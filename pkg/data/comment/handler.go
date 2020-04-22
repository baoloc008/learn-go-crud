package comment

import (
	"github.com/gin-gonic/gin"
	"learn-go-crud/pkg/crud"
)

// IRouter defines all router handle interface includes single and group router.
func RegisterService(r gin.IRouter, repo crud.Repository) {
	s := NewService(repo, newBinder())

	// get comment
	r.GET("/comment/:id", func(c *gin.Context) {
		s.Get(c, "")
	})

	r.POST("/comment", func(c *gin.Context) {
		s.Create(c, "")
	})
}
