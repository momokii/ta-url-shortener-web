package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("public/*.html")
	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "notFound.html", gin.H{})
	})

	port := "8181"

	err := r.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
