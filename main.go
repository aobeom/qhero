package main

import (
	"embed"
	"io/fs"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

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

// SeedAPI Random Number
func SeedAPI(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"message": "This is a seed",
		"data":    rand.Float64(),
	})
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
		api.GET("/seed", SeedAPI)
	}

	router.Run(":" + port)
}
