package initliza

import (
	"fmt"

	"github.com/x14n/x14n-gin-api/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() {
	global.GLogger.Info("开始数据库连接")

	mysqlConfig := global.GConfig.Mysql

	username := mysqlConfig.User
	password := mysqlConfig.Password
	host := mysqlConfig.Host
	port := mysqlConfig.Port
	dbName := mysqlConfig.Database
	timeOut := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, dbName, timeOut)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		global.GLogger.Warn("数据库连接失败")
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()

	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)

	global.G_DB = _db
	global.GLogger.Info("初始化完成")
}
