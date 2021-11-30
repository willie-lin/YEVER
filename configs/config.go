package configs

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

const harbor_base_url = "http://192.168.166.14/api/v2.0"
