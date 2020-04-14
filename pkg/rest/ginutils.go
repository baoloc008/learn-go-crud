package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetID(c *gin.Context) (int64, error) {
	id := cast.ToInt64(c.Param("id"))
	if id <= 0 {
		return 0, fmt.Errorf("missing or invalid id")
	}
	return id, nil
}
