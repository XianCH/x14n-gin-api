package initliza

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
	"github.com/x14n/x14n-gin-api/global"
)

func InitViperConfig() {
	var configFile string
	flag.StringVar(&configFile, "c", global.ConfigFile, "配置文件")
	//如果配置文件不存在
	if len(configFile) == 0 {
		panic("配置文件不存在")
	}
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置解析失败:%s\n", err))
	}
	//动态监控配置文件

	if err := v.Unmarshal(&global.GConfig); err != nil {
		panic(fmt.Errorf("配置失败:%s", err))
	}
	global.GConfig.App.ConfigFile = configFile
}
