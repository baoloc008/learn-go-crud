package crud

import "github.com/gin-gonic/gin"

type Repository interface {
	Create(data interface{}) error
	Get(ID int64) (interface{}, error)
	Delete(ID int64) error
	Update(data interface{}) error
	All() (interface{}, error)
}

// DataCreator declares interface for creating data instances type
type DataCreator interface {
	// Single: create an instance (pointer type) of the required data model, *DataType
	Single() interface{}
	// Slice: create an instance of pointer to slice of value type data model, *[]DataType
	Slice() interface{}
}

type BindType uint

const (
	BindTypeInvalid BindType = iota
	BindTypeCreate
	BindTypeUpdate
)

// DataBinder declares interface for binding rest data
type DataBinder interface {
	// Bind data and create pointer to data model
	Bind(c *gin.Context, t BindType) (interface{}, error)
}
