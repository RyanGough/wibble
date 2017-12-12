package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type NameRequest struct {
	Id int
}

func ParseNameRequest(c *gin.Context) (*NameRequest, error) {
	var r NameRequest
	if id, exists := c.GetQuery("id"); exists {
		if v, err := strconv.Atoi(id); err == nil {
			r.Id = v
		} else {
			return nil, errors.New("id should be int")
		}
	}
	return &r, nil
}
