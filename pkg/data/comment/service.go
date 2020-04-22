package comment

import (
	"github.com/gin-gonic/gin"
	"learn-go-crud/pkg/crud"
	"learn-go-crud/pkg/rest"
	"net/http"
)

type Service struct {
	repo   crud.Repository
	binder crud.DataBinder
}

func NewService(r crud.Repository, b crud.DataBinder) *Service {
	return &Service{repo: r, binder: b}
}

func (s Service) Create(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Create Comment",
	})
}

func (s Service) Get(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Get Comment",
	})
}
