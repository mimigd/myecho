package controllers

import (
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {

	param := c.QueryParam("name")

	response := map[string]interface{}{
		"message": "success",
		"code":    200,
		"data": map[string]interface{}{
			"name":  "test",
			"age":   9,
			"list1": []string{"a", "b", "c"},
			"list2": []int{1, 2, 3},
		},
		"param": param,
	}
	return c.JSON(200, response)
}
