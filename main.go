package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	// "time"

	"github.com/aobeom/minireq"

	"github.com/antchfx/htmlquery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed web/build
var build embed.FS

const (
	PreUA        = "Mozilla/5.0 (Linux; Android 7.1.1; E6533 Build/32.4.A.1.54; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.136 Mobile Safari/537.36"
	PreClient    = "jp.mdpr.mdprviewer"
	APIURL       = "https://app2-mdpr.freetls.fastly.net"
	APIModel     = "E653325 (7.1.1)"
	APIMdprApi   = "3.0.0"
	APIMdprApp   = "android:37"
	APIUserAgent = "okhttp/4.2.2"
)

// SPAIndex Import Static Files
func SPAIndex() http.FileSystem {
	fsys := fs.FS(build)
	buildStatic, _ := fs.Sub(fsys, "web/build")
	return http.FS(buildStatic)
}

// StaticHand Web Files
func StaticHand() gin.HandlerFunc {
	return func(c *gin.Context) {
		upath := c.Request.URL.Path
		if !strings.HasPrefix(upath, "/api") {
			content := SPAIndex()
			http.FileServer(content).ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

// MDPImg modelpress news images
func MDPImg(c *gin.Context) {
	mdprURL := c.Query("url")
	// Check URL
	rule := regexp.MustCompile(`https://mdpr\.jp/*`)
	result := rule.MatchString(mdprURL)
	if result {
		var imgList []string
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
		if len(imgList) != 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  1,
				"message": "MDPR Images",
				"data":    imgList,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  0,
				"message": "No Img",
				"data":    []interface{}{},
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "URL Error",
			"data":    []interface{}{},
		})
	}
}

func main() {
	port := os.Getenv("PORT")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "OPTION"}

	router := gin.Default()
	router.Use(cors.New(config))

	router.Use(StaticHand())
	api := router.Group("/api")
	{
		api.GET("/mdpr", MDPImg)
	}

	router.Run(":" + port)
}
