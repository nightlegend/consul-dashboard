package router

import (
	consulapi "consul/core/api"
	"log"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// Start a application.
func Start() {
	router := gin.New()
	// router.Use(gin.Logger())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*")
	router.GET("/index", func(c *gin.Context) {
		flag := consulapi.Put("bar", "barrierwww")
		log.Println(flag)
		result := consulapi.Get("barxx")
		log.Println(result)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8015")
}
