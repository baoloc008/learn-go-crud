package user

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

func (s Service) All(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Get All User",
	})
}

func (s Service) Create(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Create User",
	})
}

func (s Service) Get(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Get User",
	})
}

func (s Service) Remove(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Remove User",
	})
}

func (s Service) Update(c *gin.Context, tid string) {
	c.JSON(http.StatusOK, rest.ResponseData{
		Message: "Update User",
	})
}
