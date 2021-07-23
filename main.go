package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"qhero/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed web/build
var build embed.FS

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
		api.GET("/mdpr", utils.MDPImg)
	}

	router.Run(":" + port)
}
