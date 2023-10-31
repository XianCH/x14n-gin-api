package global

import (
	"github.com/x14n/x14n-gin-api/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ConfigFile = "./config.yaml"
)

var (
	GLogger *zap.Logger
	GConfig config.ServerConfig
	G_DB    *gorm.DB
)
