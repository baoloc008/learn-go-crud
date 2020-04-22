package apifx

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"learn-go-crud/pkg/crud"
	"learn-go-crud/pkg/data/comment"
	"learn-go-crud/pkg/data/user"
)

func registerHandler(r *gin.Engine, db *gorm.DB) {
	g := r.Group("/api")
	{
		user.RegisterService(g, crud.NewRepository(db, user.NewDataCreator()))
		comment.RegisterService(g, crud.NewRepository(db, comment.NewDataCreator()))
	}
}
