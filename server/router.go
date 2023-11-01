package server

import (
	"github.com/gin-gonic/gin"
	"github.com/x14n/x14n-gin-api/api/vi"
)

func NewRouter(engin *gin.Engine) {
	group := engin.Group("")
	{
		group.POST("/user/login", vi.Login)
	}
}
