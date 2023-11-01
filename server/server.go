package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/x14n/x14n-gin-api/global"
)

func getCurrentHttpServer(engin *gin.Engine) *http.Server {
	httpServer := &http.Server{
		Addr:           global.GConfig.App.Addr,
		Handler:        engin,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 >> 20,
	}
	return httpServer
}

func RunServer() {
	engine := gin.New()
	//注册公共中间件 用于处理请求处理程序中的 panic,抛出异常的时候返回http500
	engine.Use(gin.Recovery())

	httpServer := getCurrentHttpServer(engine)

	NewRouter(engine)

	printServerInfo()

	error := httpServer.ListenAndServe()
	if error != nil {
		global.GLogger.Warn(fmt.Sprintf("服务启动失败：%s", error))
		panic(error)
	}

}

func printServerInfo() {
	appConfig := global.GConfig.App
	global.GLogger.Info(fmt.Sprintf("\n【当前环境:%s 当前版本：%s 接口地址: http://%s】\n", appConfig.Env, appConfig.Version, appConfig.Addr))
	fmt.Printf("\n【当前环境:%s 当前版本：%s 接口地址: http://%s】\n", appConfig.Env, appConfig.Version, appConfig.Addr)
}
