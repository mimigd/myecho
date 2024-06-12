package main

import (
	"fmt"
	"myecho/models"
	"myecho/routes"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myecho/pkg/db"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// 初始化資料庫
	db.InitDB()
	defer db.CloseDB()

	// 自動建立資料表
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	}

	// Echo instance
	e := echo.New()

	// 初始化路由
	routes.Init(e)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// 定义一个计算请求消耗时间的中间件
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			end := time.Now()
			latency := end.Sub(start).Seconds()
			fmt.Printf("Request latency: %.6f seconds\n", latency)
			return err
		}
	})

	e.Logger.Fatal(e.Start(":8081"))
}
