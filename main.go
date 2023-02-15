package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"seeip": false,
			"ip":    "noip",
		})
	})
	router.GET("/index/ip", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"seeip": true,
			"ip":    c.ClientIP(),
		})
	})
	router.Run(":8080")
}
