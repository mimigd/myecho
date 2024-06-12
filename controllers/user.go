package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myecho/models"
)

type UserController struct {
	DB *gorm.DB
}

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

// GetUsers retrieves all users
func (uc *UserController) GetUsers(c echo.Context) error {
	var users []models.User
	err, _ := models.GetAllUsers(uc.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
