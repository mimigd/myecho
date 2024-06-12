package models

import (
	"myecho/pkg/db"
)

type User struct {
	ID       int     `gorm:"primary_key" json:"id"`
	Name     string  `gorm:"size:100;not null;comment:使用者名稱" json:"name"`
	Password string  `gorm:"size:100;not null;comment:密碼" json:"password"`
	Email    *string `gorm:"size:100;comment:信箱" json:"email"`
	Gender   *string `gorm:"size:10;comment:性別" json:"gender"`
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]User, error) {
	var users []User
	result := db.DB.Find(&users)
	return users, result.Error
}
