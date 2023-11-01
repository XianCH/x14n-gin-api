package vi

import (
	"github.com/gin-gonic/gin"
	"github.com/x14n/x14n-gin-api/global"
	"github.com/x14n/x14n-gin-api/internal/model"
	"github.com/x14n/x14n-gin-api/internal/response"
	"github.com/x14n/x14n-gin-api/internal/service"
	"go.uber.org/zap"
)

func Login(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)
	global.GLogger.Debug(user.Username + "尝试登录")

	if user.Username == "" || user.Password == "" {
		response.ErrorWithMsg(ctx, "用户名或者密码不能为空")
		return
	}

	//调用登录服务
	if err := service.UserLogin(user); err != nil {
		global.GLogger.Error("用户登录失败", zap.Any("user", user))
		response.ErrorWithMsg(ctx, "用户登录失败")
		return
	}

	//生成token
}
