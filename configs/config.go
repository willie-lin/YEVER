package configs

import (
	"github.com/spf13/viper"
	"os"
)


//var GlobalCfg *Config
//
//type Server struct {
//
//}
//
//type Mysql
//
//type Sqlite
//
//type Guacd
//
//type Sshd
//
//type Config struct {
//
//	Debug 			bool
//	Demo 			bool
//	Container       bool
//	Server 			*Server
//	Mysql			*Mysql
//	Sqlite          *Sqlite
//	ResetPassword    string
//	ResetTotp        string
//	EncryptionKey 	 string
//	EncryptPassword  []byte
//	NewEncryptionKey  string
//	Guacd			*Guacd
//	Sshd			*Sshd
//
//}


func InitConfig()  {
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

const harbor_base_url =  "http://192.168.166.14/api/v2.0"