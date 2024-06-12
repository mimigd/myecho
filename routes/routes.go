package routes

import (
	"github.com/labstack/echo/v4"
	"myecho/controllers"
)

func Init(e *echo.Echo) {
	// test
	e.GET("/test", controllers.TestAPI)

	// User routes
	e.GET("/users", controllers.GetAllUsers)
	//e.GET("/users/:id", controllers.GetUserByID)
	//e.POST("/users", controllers.CreateUser)
	//e.PUT("/users/:id", controllers.UpdateUser)
	//e.DELETE("/users/:id", controllers.DeleteUser)
	//
	//// WebSocket route
	//e.GET("/ws", handleWebSocket)
}
