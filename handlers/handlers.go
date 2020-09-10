package handlers

import (
	"fmt"
	"server/models/post"
	"server/models/user"

	"github.com/gofiber/fiber"
)

// CreatePost handles creating a post from a post request
func CreatePost(c *fiber.Ctx) {
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "documents" key:
		files := form.File["files"]
		if len(files) < 1 {
			c.Status(500).Send("a file is required")
			return
		}
		// => []*multipart.FileHeader
		storedFiles := ""
		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./static/%s", file.Filename))

			if err != nil {
				c.Status(500).Send("failed to store your file")
				return
			}
			storedFiles = file.Filename
		}

		// handle text fields
		var myPost post.Post
		// check if keys exist inside the form
		title, exists := form.Value["title"]
		if !exists {
			c.Status(500).Send("no title found in the submitted form")
			return
		}
		/* author, exists := form.Value["author"]
		if !exists {
			c.Status(500).Send("no author found in the submitted form")
			return
		} */
		myPost.Title = title[0]
		myPost.Author = "kantemir"
		myPost.ResourceName = storedFiles

		createdPost := post.CreatePost(myPost)
		c.JSON(createdPost)
	} else {
		c.Status(500).Send("failed to parse the submitted form")
	}
}

// GetPosts handles getting all the posts and sends them
func GetPosts(c *fiber.Ctx) {
	posts := post.GetPosts()
	c.JSON(posts)
}

// GetPostByID gets a single post with the id provided in the params
func GetPostByID(c *fiber.Ctx) {
	postID := c.Params("id")
	if postID == "" {
		c.Status(404).Send(fiber.Map{
			"error": "please provided a valid id",
		})
		return
	}
	post, err := post.GetPostByID(postID)
	if err != nil {
		c.Status(404).Send(fiber.Map{
			"error": "no post found with the provided id",
		})
		return
	}
	c.JSON(post)

}

// CreateUser creates trys to create a user (most likely used for signup)
func CreateUser(c *fiber.Ctx) {
	currentUser := new(user.User)
	/* currentUser.Name = "kantemir"
	currentUser.Password = "kantemir"
	currentUser.Type = 5 */
	if err := c.BodyParser(currentUser); err != nil {
		c.Status(500).Send(err)
		return
	}
	createdUser, err := user.CreateUser(*currentUser)
	if err != nil {
		c.Status(500).Send("failed to create user")
	} else {
		c.JSON(createdUser)
	}
}

// Login user
func Login(c *fiber.Ctx) {
	var currentUser user.LoginCredentials

	if err := c.BodyParser(&currentUser); err != nil {
		c.Status(500).Send(err)
		return
	}
	foundUser, err := user.GetUserByName(currentUser.Name)
	if err != nil {
		c.Status(500).Send("no user found with the given credentials")
		return
	}
	c.JSON(foundUser)

}
