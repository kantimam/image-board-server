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
		files := form.File["documents"]
		if len(files) < 1 {
			c.Status(500).Send("a file as required")
			return
		}
		// => []*multipart.FileHeader
		storedFiles := ""
		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

			if err != nil {
				c.Status(500).Send("failed to store your file")
				return
			}
			storedFiles = file.Filename
		}

		// handle text fields
		var myPost post.Post
		myPost.Title = form.Value["title"][0]
		myPost.Author = form.Value["author"][0]
		myPost.Type = "image"
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
