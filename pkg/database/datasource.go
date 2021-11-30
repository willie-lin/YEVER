package database

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/willie-lin/YEVER/pkg/database/ent"
	"github.com/willie-lin/YEVER/pkg/database/ent/migrate"
	"go.uber.org/zap"
	"log"
)

type DatabaseCfg1 struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Type     string `json:"type"`
}

type Controller struct {
	Client *ent.Client
	CTX    context.Context
	Log    *log.Logger
}

func NewClients() (*Controller, error) {

	//fmt.Println(viper.GetString("database.username"))
	return funcName()
}

func funcName() (*Controller, error) {
	var dfg = &DatabaseCfg1{
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.dbname"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Type:     viper.GetString("database.type"),
	}

	var client *ent.Client
	//controller
	var err error
	//drv, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/terminal?charset=utf8&parseTime=true")
	//client :=
	switch dfg.Type {
	case "sqlite3":
		//fmt.Println(1111111)
		client, err = ent.Open(dfg.Type, fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1", dfg.DbName))
		//fmt.Println(client)
		if err != nil {
			return &Controller{}, fmt.Errorf("failed opening connection to sqlite: %v", err)
		}
	case "mysql":
		//fmt.Println(2222222)
		client, err = ent.Open(dfg.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
			dfg.User, dfg.Password, dfg.Host, dfg.Port, dfg.DbName))
		//fmt.Println(client)
		if err != nil {
			return &Controller{}, fmt.Errorf("failed opening connection to mysql: %v", err)
		}
	case "postgres", "postgresql":
		client, err = ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			dfg.Host, dfg.Port, dfg.User, dfg.DbName, dfg.Password))
		if err != nil {
			return &Controller{}, fmt.Errorf("failed opening connection to postgres: %v", err)
		}
	default:
		return &Controller{}, fmt.Errorf("unknown database type")
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return &Controller{Client: client, CTX: context.Background(), Log: log.Default()}, nil
}

func AutoMigration1(controller *Controller) {
	log, _ := zap.NewDevelopment()
	if err := controller.Client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
		//log.Fatalf("failed creating schema resources: %v", err)
	}
}

func DebugMode1(controller *Controller, err error) {
	log, _ := zap.NewDevelopment()
	err = controller.Client.Debug().Schema.Create(
		controller.CTX,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
		//log.Fatalf("failed creating schema resources: %v", err)
	}
}
