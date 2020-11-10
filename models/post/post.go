package post

import (
	"server/database"
)

// Post is the data for a single post
type Post struct {
	ID           uint   `json:"id"`
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

type PostWithPreview struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	ResourceName string `json:"resourceName"`
	NextPost     string `json:"nextPost" gorm:"column:next_post"`
}

type PostPreview struct {
	PrevPost Preview `json:"prevPost"`
	NextPost Preview `json:"nextPost"`
}

// GetPosts returns all posts from database
func GetPosts() []Preview {
	//database.DBConn.Exec("DELETE from posts")
	var posts []Preview
	database.DBConn.Model(&Post{}).Find(&posts)
	return posts
}

/* func GetPost(condition Post) (Post, error) {
	var post Post
	response := database.DBConn.Find(condition)
	if response.Error != nil {
		return post, response.Error
	}
	return post, nil
} */

// GetPostByID trys to get a single post with a certain id from the database
func GetPostByID(id string) (Post, error) {
	var post Post
	if err := database.DBConn.First(&post, id).Error; err != nil {
		return post, err
	}
	return post, nil
}

// GetPostPreviewByID gets the next and prev post around the provided ID
func GetPostPreviewByID(id string) (PostPreview, error) {
	var postPreview PostPreview
	if err := database.DBConn.Model(&Post{}).Last(&postPreview.PrevPost, "id < ?", id).Error; err != nil {
		return postPreview, err
	}
	if err := database.DBConn.Model(&Post{}).First(&postPreview.NextPost, "id > ?", id).Error; err != nil {
		return postPreview, err
	}
	return postPreview, nil
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
