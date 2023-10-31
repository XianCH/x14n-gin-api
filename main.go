package main

import (
	"github.com/x14n/x14n-gin-api/global"
	"github.com/x14n/x14n-gin-api/initliza"
)

func main() {
	// initliza.InitLog()
	// global.GLogger.Sugar().Infof("日志写入测试:%v", strings.Repeat("hello", 6))
	// global.GLogger.Info("Info记录", zap.String("name", "张三"))
	initliza.InitLoger2("./config", "debug")
	initliza.InitGorm()
	global.GLogger.Info("server start")
	// current := time.Now()
	// fmt.Println(current)
	// formattime := current.Format("2006-01-02/15:04:05")
	// fmt.Println("current:", formattime)
}
