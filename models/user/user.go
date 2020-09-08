package user

import (
	"server/database"

	"gorm.io/gorm"
)

// User of the app
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" form:"name"`
	Password string `json:"password" form:"password"`
	Type     int    `json:"type" gorm:"default:5"  form:"type"`
}

// Preview is a simpler version of Post used to display a preview in a list
type Preview struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	ResourceName string `json:"resourceName"`
}

// GetUsers returns all posts from database
func GetUsers() []User {
	//database.DBConn.Exec("DELETE from posts")
	var user []User
	database.DBConn.Model(&User{}).Find(&user)
	return user
}

// CreateUser does what it says :3
func CreateUser(user User) User {
	database.DBConn.Create(&user)
	return user
}
