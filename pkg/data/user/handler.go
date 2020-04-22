package user

import (
	"github.com/gin-gonic/gin"
	"learn-go-crud/pkg/crud"
)

// IRouter defines all router handle interface includes single and group router.
func RegisterService(r gin.IRouter, repo crud.Repository) {
	s := NewService(repo, newBinder())

	// get all user
	r.GET("/users", func(c *gin.Context) {
		s.All(c, "")
	})

	// create a user
	r.POST("/user", func(c *gin.Context) {
		s.Create(c, "")
	})

	// get user with id
	r.GET("/user/:id", func(c *gin.Context) {
		s.Get(c, "")
	})

	// delete user with id, can use r.DELETE
	r.DELETE("/user/:id", func(c *gin.Context) {
		s.Remove(c, "")
	})

	// update user with id, can use r.PUT
	r.PUT("/user/:id", func(c *gin.Context) {
		s.Update(c, "")
	})
}
