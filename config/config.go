package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	cfg Cfg
)

type Cfg struct {
	*MysqlCfg
	*ServerCfg
	*ZapCfg
}

func GetMysqlConfig() *MysqlCfg {
	return cfg.MysqlCfg
}

func GetServerConfig() *ServerCfg {
	return cfg.ServerCfg
}

func GetZapConfig() *ZapCfg {
	return cfg.ZapCfg
}

func LoadConfig(filePath string) {
	viper.SetConfigFile(filePath) // name of config file (without extension)
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	parseConfig()
}

func parseConfig() {
	err := viper.Sub("mysql").Unmarshal(&cfg.MysqlCfg)
	if err != nil {
		panic(err)
	}
	err = viper.Sub("server").Unmarshal(&cfg.ServerCfg)
	if err != nil {
		panic(err)
	}
	err = viper.Sub("zap").Unmarshal(&cfg.ZapCfg)
	if err != nil {
		panic(err)
	}
}
