package loader

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

const configRelativePath = "/config"

// Config 配置文件
var Config *viper.Viper

// Env 当前环境 prod | pre | test | dev | local
const (
	EnvProd  = "prod"
	EnvPre   = "pre"
	EnvTest  = "test"
	EnvDev   = "dev"
	EnvLocal = "local"
)

// loadEnvConfig 读取环境相关配置
func LoadEnvConfig(path string) {
	var configPath string
	if path != "" {
		configPath = path + configRelativePath
	} else {
		configPath = GetExecPath() + configRelativePath
	}
	Config = viper.New()
	Config.AddConfigPath(configPath)

	switch viper.GetString("env") {
	case EnvProd:
		Config.SetConfigName(EnvProd)
	case EnvPre:
		Config.SetConfigName(EnvPre)
	case EnvTest:
		Config.SetConfigName(EnvTest)
	case EnvDev:
		Config.SetConfigName(EnvDev)
	case EnvLocal:
		Config.SetConfigName(EnvLocal)
	}

	err := Config.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err)
	}
}

// 获取执行文件的目录
func GetExecPath() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return ex
}
