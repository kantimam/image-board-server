package post

import (
	"server/database"

	"gorm.io/gorm"
)

// Post is the data for a single post
type Post struct {
	gorm.Model
	Title        string `json:"title"`
	Author       string `json:"author"`
	ResourceName string `json:"resourceName"`
}

// Preview is a simpler version of Post used to display a preview in a list
type Preview struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	ResourceName string `json:"resourceName"`
}

// GetPosts returns all posts from database
func GetPosts() []Preview {
	//database.DBConn.Exec("DELETE from posts")
	var posts []Preview
	database.DBConn.Model(&Post{}).Find(&posts)
	return posts
}

// CreatePost does what it says :3
func CreatePost(post Post) Post {
	// TODO extract data outside of here instead of using the request
	/* var post Post
	post.Title = "first post"
	post.Author = "1414"
	post.ResourceName = "cat"
	post.Type = "image" */

	database.DBConn.Create(&post)
	return post
}
