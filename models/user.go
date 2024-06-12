package models

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `gorm:"size:100" json:"name"`
	Password string `gorm:"size:100" json:"password"`
	Email    string `gorm:"size:100" json:"email"`
}

// GetAllUsers retrieves all users
func GetAllUsers(db *gorm.DB) ([]*User, error) {
	var users []*User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
