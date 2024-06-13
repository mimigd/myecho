package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"myecho/def"
	"myecho/models"
	"myecho/pkg/rdb"
)

var testValue = 0

func TestAPI(c echo.Context) error {

	param := c.QueryParam("name")

	testValue++

	ping := rdb.RDB.Ping(c.Request().Context())

	response := map[string]interface{}{
		"message": "success",
		"code":    def.RETURN_CODE_SUCCESS,
		"data": map[string]interface{}{
			"name":  "test",
			"age":   9,
			"list1": []string{"a", "b", "c"},
			"list2": []int{1, 2, 3},
		},
		"param":   param,
		"test":    testValue,
		"context": c.Request().Context(),
		"ping":    ping.Val(),
	}
	return c.JSON(200, response)
}

// GetUsers retrieves all users
func GetAllUsers(c echo.Context) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
