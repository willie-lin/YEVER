package database

import (
	"context"
	"entgo.io/ent/entc/integration/ent/migrate"
	"entgo.io/ent/examples/fs/ent"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type DatabaseCfg struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Type     string `json:"type"`
}

func NewClient() (*ent.Client, error) {

	//fmt.Println(viper.GetString("database.username"))
	var dfg = &DatabaseCfg{
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
			return client, fmt.Errorf("failed opening connection to sqlite: %v", err)
		}
	case "mysql":
		//fmt.Println(2222222)
		client, err = ent.Open(dfg.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
			dfg.User, dfg.Password, dfg.Host, dfg.Port, dfg.DbName))
		//fmt.Println(client)
		if err != nil {
			return client, fmt.Errorf("failed opening connection to mysql: %v", err)
		}
	case "postgres", "postgresql":
		client, err = ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			dfg.Host, dfg.Port, dfg.User, dfg.DbName, dfg.Password))
		if err != nil {
			return client, fmt.Errorf("failed opening connection to postgres: %v", err)
		}
	default:
		return client, fmt.Errorf("unknown database type")
	}
	return client, err
}

func AutoMigration(client *ent.Client, ctx context.Context) {
	log, _ := zap.NewDevelopment()
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
		//log.Fatalf("failed creating schema resources: %v", err)
	}
}

func DebugMode(err error, client *ent.Client, ctx context.Context) {
	log, _ := zap.NewDevelopment()
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
		//log.Fatalf("failed creating schema resources: %v", err)
	}
}
