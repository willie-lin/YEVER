package main

import (
	"context"
	"fmt"
	"github.com/bykof/gostradamus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/willie-lin/YEVER/configs"
	"github.com/willie-lin/YEVER/pkg/database"
	"github.com/willie-lin/YEVER/pkg/handler"
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
	//fmt.Println(client)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
		return
	}

	defer client.Close()

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

	controller := &handler.Controller{Client: client}

	// Server

	// Routes
	e.GET("/", hello)
	e.GET("/users", handler.GetAllUser(client))
	e.POST("/user", handler.CreateUser(client))
	e.GET("/user/id", handler.FindUserByName(client))
	e.DELETE("/user", handler.DeleteUser(client))
	e.PUT("/user/name", handler.UpdateUserByName(client))
	e.PUT("/user/id", handler.UpdateUserByID(client))

	//e.PATCH("/user", handler.UpdateUser(client))
	//e.POST("/user", )
	e.POST("/user1", controller.InsertComment)
	e.Static("/", "static/")

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
