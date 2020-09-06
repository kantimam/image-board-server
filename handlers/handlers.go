package handlers

import (
	"fmt"
	"server/models/post"

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
