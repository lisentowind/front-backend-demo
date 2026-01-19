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
		Msg:  "ping",
		Data: gin.H{"server_time": utils.GetChinaTime()},
	}
	c.JSON(http.StatusOK, resp)
}

func (h *HelloController) HelloTable(c *gin.Context) {

	list := make([]utils.User, 0)

	for i := 1; i <= 10; i++ {
		list = append(list, utils.RandomUser(i))
	}

	tableData := model.TableData{
		List:  list,
		Page:  1,
		Size:  10,
		Total: len(list),
	}

	c.JSON(http.StatusOK, model.TableResponse{
		Code: 200,
		Msg:  "hello table",
		Data: tableData,
	})
}
