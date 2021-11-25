package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/willie-lin/trivy-scan-images/configs"

	"net/http"
)

func main() {

	//log, _ := zap.NewDevelopment()

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

	// 连接 数据库

	//client, err := database.NewClient()
	//if err != nil {
	//	log.Fatal("opening ent client", zap.Error(err))
	//	return
	//}
	//dateTime := gostradamus.Now()
	//fmt.Println(dateTime)
	//
	//defer client.Close()
	//ctx := context.Background()
	//
	////autoMigration := database.AutoMigration
	//autoMigration := database.AutoMigration
	//autoMigration(client, ctx)
	//
	////debugMode := database.DebugMode
	//debugMode := database.DebugMode
	//
	//debugMode(err, client, ctx)
	//
	//fmt.Println(viper.GetString("database.password"))


	// Server
	// Routes
	e.GET("/", hello)
	e.GET("/user", configs.TestGoHarborClient())

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
