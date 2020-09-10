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

// BasicUser a simple User object
type BasicUser struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// LoginCredentials are the fields that are required to login
type LoginCredentials struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
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
	response := database.DBConn.First(&user, id)
	if response.Error != nil {
		return user, response.Error
	}
	return user, nil
}

// GetUserByID finds a user by its ID and returns the entire user data (should only be used to check login)
/* func GetUserByID(id uint) (BasicUser, error) {
	var user BasicUser
	response := database.DBConn.Model(&User{}).First(id)
	if response.Error != nil {
		return user, response.Error
	}
	return user, nil
} */

// GetUserByName trys to find a user with the given username (will be used for login)
func GetUserByName(name string) (User, error) {
	var user User
	response := database.DBConn.Where(&User{Name: name}).First(&user)
	if response.Error != nil {
		return user, response.Error
	}
	return user, nil
}

// CreateUser does what it says :3
func CreateUser(user User) (User, error) {
	response := database.DBConn.Create(&user)
	if response.Error != nil {
		return user, response.Error
	}
	return user, nil
}
