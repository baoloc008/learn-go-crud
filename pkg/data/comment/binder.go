package comment

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
	case crud.BindTypeRead:
		return b.bindGet(c)
	case crud.BindTypeInvalid:
		return nil, fmt.Errorf("invalid BindType")
	default:
		return nil, fmt.Errorf("unknown BindType")
	}
}

func (b binder) bindCreate(c *gin.Context) (*models.Comment, error) {
	data := rest.CreateCommentRequestData{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}

	return &models.Comment{
		Content: fmt.Sprintf("%s: %s", data.UserName, data.Message),
	}, nil
}

func (b binder) bindGet(c *gin.Context) (*models.Comment, error) {
	return &models.Comment{
		BaseModel: models.BaseModel{ID: 1},
	}, nil
}
