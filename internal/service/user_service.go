package service

import (
	"github.com/x14n/x14n-gin-api/global"
	"github.com/x14n/x14n-gin-api/internal/model"
)

type UserDao struct {
	Uid uint
}

func UserLogin(userModel model.User) error {
	result := global.G_DB.Where("username = ? and password = ?", userModel.Username, userModel.Password).First(userModel)
	if result.Error != nil {
		return result.Error
	}

	return result.Error
}

// todo
func (u *UserDao) FindUserById() (*model.User, error) {
	var user model.User
	result := global.G_DB.Where("id = ?", u.Uid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	//查询用户信息
	userInfo := model.UserInfo{}
	result = global.G_DB.Where("uid = ?", u.Uid).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	// return user, nil
	return nil, nil
}
