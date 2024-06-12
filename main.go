package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myecho/controllers"
	"time"
)

func main() {
	//startTime := time.Now().UnixNano()

	e := echo.New()

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

	e.GET("/users", controllers.GetUser)
	e.Logger.Fatal(e.Start(":8081"))

	// 計算消耗時間
	//endTime := time.Now().UnixNano()
	//duration := float64(endTime-startTime) / 1000000000
	//fmt.Printf("消耗時間(秒): %.6f\n", duration)
}
