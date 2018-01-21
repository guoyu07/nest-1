package main

import (
	"github.com/wolfogre/nest/internal/spider"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/wolfogre/nest/internal/service/util/timeformat"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	spider.StartDaemon()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("tmpl/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "loupan-heat.tmpl", gin.H{
			"key": "d5c59bd4a0caf3f9de7372cf2f3d7bb2", // key 绑定域名的，泄露也没关系
			"date": timeformat.FormatDateNow(),
		})
	})
	router.Static("/static", "./static/")
	router.Run(":80")
}
