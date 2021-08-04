package lib

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryAsInt(c *gin.Context, name string) (int, error) {
	val := c.Query(name)
	if val == "" {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}
	return strconv.Atoi(val)
}

func GetPostFormAsInt(c *gin.Context, name string) (int, error) {
	val := c.PostForm(name)
	if val == "" {
		return 0, errors.New(name + " post form parameter value is empty or not specified")
	}
	return strconv.Atoi(val)
}
