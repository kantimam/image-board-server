package handlers

import (
	"fmt"
	"server/models/post"
	"server/models/user"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

// CreatePost handles creating a post from a post request
func CreatePost(c *fiber.Ctx) error {
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "documents" key:
		files := form.File["files"]
		if len(files) < 1 {
			return c.Status(500).SendString("a file is required")

		}
		// => []*multipart.FileHeader
		storedFiles := ""
		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			targetFilePath := fmt.Sprintf("./static/%s", file.Filename)
			// Save the files to disk:
			err := c.SaveFile(file, targetFilePath)

			if err != nil {
				return c.Status(500).SendString("failed to store your file")

			}
			storedFiles = file.Filename
			go util.StoreThumbnails(targetFilePath, storedFiles)

		}

		// handle text fields
		var myPost post.Post
		// check if keys exist inside the form
		title, exists := form.Value["title"]
		if !exists {
			return c.Status(500).SendString("no title found in the submitted form")

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
		return c.JSON(createdPost)
	} else {
		return c.Status(500).SendString("failed to parse the submitted form")
	}
}

// GetPosts handles getting all the posts and sends them
func GetPosts(c *fiber.Ctx) error {
	posts := post.GetPosts()
	return c.JSON(posts)
}

// GetPostByID gets a single post with the id provided in the params
func GetPostByID(c *fiber.Ctx) error {
	postID := c.Params("id")
	if postID == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": "please provided a valid id",
		})
	}
	post, err := post.GetPostByID(postID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "no post found with the provided id",
		})
	}
	return c.JSON(post)
}

func GetPostPreview(c *fiber.Ctx) error {
	postID := c.Params("id")
	if postID == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": "please provided a valid id",
		})

	}
	post, err := post.GetPostPreviewByID(postID)
	if err != nil {
		fmt.Println(err)
		return c.Status(404).JSON(fiber.Map{
			"error": "no post found with the provided id",
		})

	}
	return c.JSON(post)
}

// CreateUser creates trys to create a user (most likely used for signup)
func CreateUser(c *fiber.Ctx) error {
	currentUser := new(user.User)
	/* currentUser.Name = "kantemir"
	currentUser.Password = "kantemir"
	currentUser.Type = 5 */
	if err := c.BodyParser(currentUser); err != nil {
		return c.Status(500).JSON(err)

	}
	createdUser, err := user.CreateUser(*currentUser)
	if err != nil {
		return c.Status(500).SendString("failed to create user")
	} else {
		return c.JSON(createdUser)
	}
}

// Login user
func Login(c *fiber.Ctx) error {
	var currentUser user.LoginCredentials

	if err := c.BodyParser(&currentUser); err != nil {
		return c.Status(500).JSON(err)

	}
	foundUser, err := user.GetUserByName(currentUser.Name)
	if err != nil {
		return c.Status(500).SendString("no user found with the given credentials")
	}
	return c.JSON(foundUser)

}
