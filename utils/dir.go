package utils

import (
	"os"

	"github.com/x14n/x14n-gin-api/global"
)

func DirExit(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateDir(path string) error {
	isExit, err := DirExit(path)
	if err != nil {
		return err
	}
	if !isExit {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			global.GLogger.Sugar().Debugf("创建文件[%s]目录失败:%s", path, err)
		}
	}
	return err
}
