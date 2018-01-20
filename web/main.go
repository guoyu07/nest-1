package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("tmpl/*")
	router.GET("/loupan-heat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "loupan-heat.tmpl", gin.H{
			"key": "d5c59bd4a0caf3f9de7372cf2f3d7bb2",
		})
	})
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.tmpl", gin.H{
			"key": "d5c59bd4a0caf3f9de7372cf2f3d7bb2",
		})
	})
	router.Run(":80")
}
