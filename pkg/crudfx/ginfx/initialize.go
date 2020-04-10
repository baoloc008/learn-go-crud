package ginfx

import "github.com/gin-gonic/gin"

func provideGin() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.Writer.Write([]byte("pong"))
	})

	return r
}
