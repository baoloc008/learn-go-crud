package crud

import (
	"github.com/gin-gonic/gin"
	"learn-go-crud/pkg/rest"
	"net/http"
)

type Service struct {
	repo   Repository
	binder DataBinder
}

func NewService(r Repository, b DataBinder) *Service {
	return &Service{repo: r, binder: b}
}

func (s Service) All(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "All",
	})
}

func (s Service) Create(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Create",
	})
}

func (s Service) Get(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Get",
	})
}

func (s Service) Remove(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Remove",
	})
}

func (s Service) Update(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Update",
	})
}
