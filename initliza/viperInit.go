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
	if err := v.ReadConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败:%s\n", err))
	}

}
