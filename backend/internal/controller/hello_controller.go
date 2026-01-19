package controller

import (
	"backend/internal/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (h *HelloController) Ping(c *gin.Context) {
	resp := model.Response{
		Code: 200,
		Msg:  "pong",
		Data: gin.H{"server_time": utils.GetChinaTime()},
	}
	c.JSON(http.StatusOK, resp)
}
