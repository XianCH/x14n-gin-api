package initliza

import (
	"path"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/x14n/x14n-gin-api/global"
	"github.com/x14n/x14n-gin-api/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	outJson = "json"
)

func InitLog() {

	logConfig := global.GConfig.Log
	if exit, _ := utils.DirExit(logConfig.Path); !exit {
		err := utils.CreateDir(logConfig.Path)
		if err != nil {
			panic("初始化日志失败")
		}
	}
	//设置输出格式
	var encoder zapcore.Encoder
	if logConfig.OutFormat == outJson {
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	//设置日志切割
	writeSyncer := zapcore.AddSync(getLumberjackWriteSyncer())
	// 创建NewCore
	zapCore := zapcore.NewCore(encoder, writeSyncer, getLevel())
	//创建logger
	logger := zap.New(zapCore)
	defer logger.Sync()
	global.GLogger = logger
}

func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime, // 自定义输出时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return config
}

func getLevel() zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if level, ok := levelMap[global.GConfig.Log.Level]; ok {
		return level
	}
	// 默认info级别
	return zapcore.InfoLevel
}

func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// 获取文件切割和归档配置信息
func getLumberjackWriteSyncer() zapcore.WriteSyncer {
	lumberjackConfig := global.GConfig.Log.LumberJack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   getLogFile(),                //日志文件
		MaxSize:    lumberjackConfig.MaxSize,    //单文件最大容量(单位MB)
		MaxBackups: lumberjackConfig.MaxBackups, //保留旧文件的最大数量
		MaxAge:     lumberjackConfig.MaxAge,     // 旧文件最多保存几天
		Compress:   lumberjackConfig.Compress,   // 是否压缩/归档旧文件
	}
	// 设置日志文件切割
	return zapcore.AddSync(lumberjackLogger)
}

func getLogFile() string {
	fileFormat := time.Now().Format(global.GConfig.Log.FileFormat)
	fileName := strings.Join([]string{
		global.GConfig.Log.FilePrefix,
		fileFormat,
		"log"}, ".")
	return path.Join(global.GConfig.Log.Path, fileName)
}
