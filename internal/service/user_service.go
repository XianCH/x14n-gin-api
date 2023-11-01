package service

import (
	"github.com/x14n/x14n-gin-api/global"
	"github.com/x14n/x14n-gin-api/internal/model"
)

func UserLogin(userModel model.User) error {
	result := global.G_DB.Where("username = ? and password = ?", userModel.Username, userModel.Password).First(userModel)
	if result.Error != nil {
		return result.Error
	}

	return result.Error
}
