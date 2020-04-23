package user

import (
	"github.com/gin-gonic/gin"
	"learn-go-crud/pkg/crud"
	"learn-go-crud/pkg/logger"
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
	data, err := s.repo.All()
	if err != nil {
		logger.Errorw("Error while get all roles", "error", err)
		c.JSON(http.StatusOK, rest.ResponseData{
			Code:    rest.CannotGetErrorCode,
			Message: "Get All User Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rest.ResponseData{
		Code:    rest.SuccessCode,
		Message: "Get All User",
		Data:    data,
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
