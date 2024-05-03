package config

import (
	"github.com/spf13/viper"
	"tron-tools/pkg/go-logger"
)

var Conf *Config

func Initialize() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./yaml")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("ReadInConfig error:", err)
	}

	Conf = &Config{
		LogCfg: viper.GetString("log"),
		Console: &Console{
			Name:    viper.GetString("console.name"),
			Version: viper.GetString("console.version"),
			Port:    viper.GetString("console.port"),
		},
		Tron: &Tron{
			NodeUrl:      viper.GetString("tron.nodeurl"),
			ContractAddr: viper.GetString("tron.contractAddr"),
		},
	}

	// 初始化日志配置
	logger.SetLogger("./yaml/log.json")
	//成功加载配置
	logger.Info("Successful configuration load")
}
