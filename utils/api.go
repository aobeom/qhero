package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/aobeom/minireq"
	"github.com/gin-gonic/gin"
)

const (
	PreUA        = "Mozilla/5.0 (Linux; Android 7.1.1; E6533 Build/32.4.A.1.54; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.136 Mobile Safari/537.36"
	PreClient    = "jp.mdpr.mdprviewer"
	APIURL       = "https://app2-mdpr.freetls.fastly.net"
	APIModel     = "E653325 (7.1.1)"
	APIMdprApi   = "3.0.0"
	APIMdprApp   = "android:37"
	APIUserAgent = "okhttp/4.2.2"
)

// MediaFromDB 读取 media 数据
func MediaFromDB(url string) (data map[string]interface{}) {
	sql := "SELECT url, source FROM mdpr WHERE url=$1"
	row := Psql.QueryOne(sql, url)

	var murl string
	var msource QMediaArray

	row.Scan(&murl, &msource)
	if murl != "" {
		data = map[string]interface{}{
			"url":    murl,
			"source": msource,
		}
	}
	return
}

// Media2DB 保存 media 的数据
func Media2DB(url string, sources interface{}) {
	sql := "INSERT INTO mdpr (url,source) VALUES ($1,$2)"

	var newSources QMediaArray
	for _, source := range sources.([]string) {
		newSources = append(newSources, source)
	}

	Psql.Exec(sql, url, newSources)
}

func filterMdpr(mdprURL string) (imgList []string) {
	urlSplit := strings.Split(mdprURL, "/")
	articleID := urlSplit[len(urlSplit)-1]
	preURL := fmt.Sprintf("https://app2-mdpr.freetls.fastly.net/articles/detail/%s", articleID)

	requests := minireq.Requests()
	// Pre Data
	preHeader := minireq.Headers{
		"User-Agent":       PreUA,
		"X-Requested-With": PreClient,
	}
	preRes := requests.Get(preURL, preHeader)
	doc, _ := htmlquery.Parse(strings.NewReader(string(preRes.RawData())))
	nodes := htmlquery.Find(doc, `//div[@class="p-articleBody"]/a`)
	apiData := htmlquery.SelectAttr(nodes[0], "data-mdprapp-option")
	jsonRaw, _ := url.QueryUnescape(apiData)
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(jsonRaw), &jsonData)
	imgURL := APIURL + jsonData["url"].(string)
	// Img Data
	apiHeader := minireq.Headers{
		"model":      APIModel,
		"mdpr-api":   APIMdprApi,
		"mdpr-app":   APIMdprApp,
		"User-Agent": APIUserAgent,
	}
	apiRes := requests.Get(imgURL, apiHeader)
	var imgJson map[string]interface{}
	json.Unmarshal(apiRes.RawData(), &imgJson)
	imgData := imgJson["list"]
	for _, img := range imgData.([]interface{}) {
		iData := img.(map[string]interface{})
		imgList = append(imgList, iData["url"].(string))
	}
	return imgList
}

// MDPImg modelpress news images
func MDPImg(c *gin.Context) {
	mdprURL := c.Query("url")
	// Check URL
	rule := regexp.MustCompile(`https://mdpr\.jp/*`)
	result := rule.MatchString(mdprURL)
	if result {
		sources := MediaFromDB(mdprURL)
		if len(sources) != 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  1,
				"message": "MDPR Images",
				"data":    sources["source"],
				"cache":   true,
			})
		} else {
			imgList := filterMdpr(mdprURL)
			if len(imgList) != 0 {
				Media2DB(mdprURL, imgList)
				c.JSON(http.StatusOK, gin.H{
					"status":  1,
					"message": "MDPR Images",
					"data":    imgList,
					"cache":   false,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  0,
					"message": "No Img",
					"data":    []interface{}{},
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "URL Error",
			"data":    []interface{}{},
		})
	}
}
