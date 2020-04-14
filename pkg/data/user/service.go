package user

import (
	"github.com/gin-gonic/gin"
	"learn-go-crud/pkg/crud"
)

// IRouter defines all router handle interface includes single and group router.
func RegisterService(r gin.IRouter, repo crud.Repository) {
	s := crud.NewService(repo, newBinder())
	r.GET("/users", func(c *gin.Context) {
		s.All(c, "")
	})
	r.POST("/user", func(c *gin.Context) {
		s.Create(c, "")
	})
	r.GET("/user/:id", func(c *gin.Context) {
		s.Get(c, "")
	})
	r.DELETE("/user/:id", func(c *gin.Context) {
		s.Remove(c, "")
	})
	r.PUT("/user/:id", func(c *gin.Context) {
		s.Update(c, "")
	})
}
