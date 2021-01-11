package loader

import (
	"log"

	"github.com/spf13/viper"
)

// 读取配置文件
func LoadConfig(path string) {
	bindEnv()

	if viper.GetString("env") == "" {
		log.Fatal("没有指定[MICRO_SERVICE_ENV]环境变量,usage export MICRO_SERVICE_ENV=dev or usage export MICRO_SERVICE_ENV=prod")
	}
	log.Printf("当前运行环境 %s", viper.GetString("env"))

	LoadEnvConfig(path)
}

func bindEnv() {
	viper.SetEnvPrefix("micro_service")
	viper.BindEnv("env") // 区分线上线下的环境变量 dev | prod
}
