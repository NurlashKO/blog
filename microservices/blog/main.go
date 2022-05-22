package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("/templates/index.tmpl")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
		//c.JSON(200, gin.H{
		//	"message": "Hello World!",
		//	"source code": "https://github.com/NurlashKO/blog",
		//	"docs": "https://docs.nurlashko.de",
		//	"grafana": "https://grafana.nurlashko.de",
		//	"prometheus": "https://prometheus.nurlashko.de",
		//})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}