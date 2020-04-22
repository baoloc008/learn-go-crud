package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go-crud/pkg/crud"
	"learn-go-crud/pkg/data/models"
	"learn-go-crud/pkg/rest"
)

type binder struct{}

func newBinder() crud.DataBinder {
	return &binder{}
}

/*-- Implement crud.DataBinder --*/

func (b binder) Bind(c *gin.Context, t crud.BindType) (interface{}, error) {
	switch t {
	case crud.BindTypeCreate:
		return b.bindCreate(c)
	case crud.BindTypeUpdate:
		return b.bindUpdate(c)
	case crud.BindTypeInvalid:
		return nil, fmt.Errorf("invalid BindType")
	default:
		return nil, fmt.Errorf("unknown BindType")
	}
}

func (b binder) bindCreate(c *gin.Context) (*models.User, error) {
	data := rest.CreateUserRequestData{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}

	return &models.User{
		DisplayName: data.DisplayName,
		Email:       data.Email,
		UserName:    data.UserName,
	}, nil
}

func (b binder) bindUpdate(c *gin.Context) (*models.User, error) {
	id, err := rest.GetID(c)
	if err != nil {
		return nil, err
	}
	data := rest.UpdateUserRequestData{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}
	return &models.User{
		BaseModel:   models.BaseModel{ID: id},
		DisplayName: data.DisplayName,
		Email:       data.Email,
		UserName:    data.UserName,
	}, nil
}
