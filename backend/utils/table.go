package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageQuery struct {
	Page int `form:"pageNum" binding:"required,min=1"`
	Size int `form:"pageSize" binding:"required,min=1"`
}

func GetPage(c *gin.Context) PageQuery {
	page, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	if size > 100 {
		size = 100
	}

	return PageQuery{Page: page, Size: size}
}
