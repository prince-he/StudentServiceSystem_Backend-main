package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteUserRequest struct {
	Username string `json:"username" binding:"required"`
}

func DeleteUser(c *gin.Context) {
	var req DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}

	info, err := service.GetUserByUserID(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取管理员信息失败", zap.Error(err))
		utils.JsonFail(c, 200512, "获取管理员信息失败")
		return
	}

	if info.UserType != 3 {
		utils.JsonFail(c, 200511, "当前用户不是超级管理员")
		return
	}

	// 检查用户是否存在
	user, err := service.GetUserByUsername(req.Username)
	if err != nil {
		utils.JsonFail(c, 200514, "用户不存在")
		return
	}

	// 调用服务层删除用户
	service.DeleteUser(user.ID)

	utils.JsonSuccess(c, nil)
}
