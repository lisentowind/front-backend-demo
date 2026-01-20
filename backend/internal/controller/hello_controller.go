package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-backend-demo/config"
	"go-backend-demo/internal/model"
	"go-backend-demo/utils"
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

// 删除用户（支持单个和批量）
func (h *HelloController) DeleteUsers(c *gin.Context) {
	var req struct {
		Ids []int `json:"ids" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误，请提供要删除的用户ID列表",
			Data: nil,
		})
		return
	}

	// 执行批量删除
	result := config.DB.Where("id IN ?", req.Ids).Delete(&model.User{})
	if result.Error != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "删除失败: " + result.Error.Error(),
			Data: nil,
		})
		return
	}

	// 检查是否有记录被删除
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, model.Response{
			Code: http.StatusNotFound,
			Msg:  "未找到要删除的用户",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  "删除成功",
		Data: gin.H{"deletedCount": result.RowsAffected},
	})
}

func (h *HelloController) GetAllUsers(c *gin.Context) {
	var pageQuery utils.PageQuery
	if err := c.ShouldBindQuery(&pageQuery); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "分页参数错误",
			Data: nil})
		return
	}

	var users []model.User
	var total int64

	db := config.DB.Model(&model.User{})

	if err := db.Count(&total).Error; err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "获取用户总数失败",
			Data: nil})
		return
	}

	offset := (pageQuery.Page - 1) * pageQuery.Size

	if err := db.
		Order("id DESC").
		Offset(offset).
		Limit(pageQuery.Size).
		Find(&users).Error; err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "获取用户总数失败",
			Data: nil})
		return
	}

	c.JSON(http.StatusOK,
		model.Response{
			Code: http.StatusOK,
			Msg:  "获取用户总数成功",
			Data: model.TableData{
				List:  users,
				Page:  pageQuery.Page,
				Size:  pageQuery.Size,
				Total: int(total),
			}},
	)
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
