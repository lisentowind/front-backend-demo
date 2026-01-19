package controller

import (
	"backend/config"
	"backend/internal/model"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (h *HelloController) Ping(c *gin.Context) {
	resp := model.Response{
		Code: http.StatusOK,
		Msg:  "ping",
		Data: gin.H{"server_time": utils.GetChinaTime()},
	}
	c.JSON(http.StatusOK, resp)
}

// 添加用户
func (h *HelloController) AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp := model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "参数错误",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// 设置创建时间
	user.CreateTime = utils.GetChinaTime()

	// 保存到数据库
	if err := config.DB.Create(&user).Error; err != nil {
		resp := model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "添加失败: " + err.Error(),
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := model.Response{
		Code: http.StatusOK,
		Msg:  "添加成功",
		Data: user,
	}
	c.JSON(http.StatusOK, resp)
}

// 根据ID获取用户信息
func (h *HelloController) GetUserById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, model.Response{
			Code: http.StatusBadRequest,
			Msg:  "用户ID不能为空",
			Data: nil,
		})
		return
	}

	var user model.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: http.StatusNotFound,
			Msg:  "用户不存在",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  "获取用户信息成功",
		Data: user,
	})
}

func (h *HelloController) GetAllUsers(c *gin.Context) {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: http.StatusOK,
			Msg:  "获取用户列表失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  "获取用户列表成功",
		Data: users,
	})
}

func (h *HelloController) HelloTable(c *gin.Context) {

	list := make([]model.User, 0)

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
		Code: http.StatusOK,
		Msg:  "hello table",
		Data: tableData,
	})
}
