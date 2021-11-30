package main

import (
	"context"
	"fmt"
	"github.com/bykof/gostradamus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/willie-lin/YEVER/configs"
	"github.com/willie-lin/YEVER/pkg/database"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	log, _ := zap.NewDevelopment()

	e := echo.New()
	// Middleware
	//e.Use(apmechov4.Middleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//连接 数据库

	client, err := database.NewClient()
	//client, err = database.NewClients()

	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))

	}

	//controller := database.Controller{Client: client, CTX: context.Background(),}

	//fmt.Println(client)
	dateTime := gostradamus.Now()
	fmt.Println(dateTime)

	//defer client.Close()
	ctx := context.Background()

	//autoMigration := database.AutoMigration
	autoMigration := database.AutoMigration
	autoMigration(client, ctx)
	//
	//debugMode := database.DebugMode
	debugMode := database.DebugMode
	//
	debugMode(err, client, ctx)
	//
	//fmt.Println(viper.GetString("database.password"))

	// Server
	// Routes
	e.GET("/", hello)
	//e.GET("/user", configs.TestGoHarborClient())
	//e.POST("/user", )

	// Start server
	e.Logger.Fatal(e.Start(":2022"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func init() {
	configs.InitConfig()
}
