package user

import (
	"learn-go-crud/pkg/crud"
	"learn-go-crud/pkg/data/models"
)

type dataCreator struct{}

func NewDataCreator() crud.DataCreator {
	return &dataCreator{}
}

/*-- Implement crud.DataCreator --*/

func (dataCreator) Single() interface{} {
	return &models.User{}
}

func (dataCreator) Slice() interface{} {
	return &[]models.User{}
}
