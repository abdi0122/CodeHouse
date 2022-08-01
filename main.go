package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./codehouse-vue/dist", false)))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Go and Gin")
	})
	r.Run(":8090")
}