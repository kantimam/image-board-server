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
}

// lets not deal with admins and moderators for now :)
/* type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" form:"name"`
	Password string `json:"password" form:"password"`
	Type     int    `json:"type" gorm:"default:5"  form:"type"`
}
*/
type BasicUser struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetUsers returns all posts from database
func GetUsers() []User {
	//database.DBConn.Exec("DELETE from posts")
	var user []User
	database.DBConn.Model(&User{}).Find(&user)
	return user
}

// GetBasicUserByID finds a user by its ID and returns a simple user object without password obiously (should be use to get user data)
func GetBasicUserByID(id uint) (User, error) {
	var user User
	response := database.DBConn.Model(&User{}).First(id)
	if response.Error != nil {
		return user, response.Error
	}
	return user, nil
}

// GetUserByID finds a user by its ID and returns the entire user data (should only be used to check login)
func GetUserByID(id uint) (BasicUser, error) {
	var user BasicUser
	response := database.DBConn.Model(&User{}).First(id)
	if response.Error != nil {
		return user, response.Error
	}
	return user, nil
}

// CreateUser does what it says :3
func CreateUser(user User) User {
	database.DBConn.Create(&user)
	return user
}
