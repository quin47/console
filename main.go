package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	router := gin.Default()

	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	//server front resource
	router.StaticFile("/web","./public/index.html")

	router.Static("/static","./public/static")

	//api
	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//redirect to web just for now
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web")
	})


	router.Run(":8000")
}

