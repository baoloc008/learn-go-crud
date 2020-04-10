package ginfx

import "github.com/gin-gonic/gin"

func provideGin() *gin.Engine {
	r := gin.Default()
	return r
}
