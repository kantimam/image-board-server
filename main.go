package main

import (
	"fmt"
	"log"
	"os"
	"server/database"
	"server/models/post"

	"github.com/gofiber/fiber"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Settings for the server setup
type Settings struct {
	ServerPort        int    `default:"5600"`
	StaticFilesFolder string `default:"./static"`
}

func initDB() {
	var err error
	dsn := "host=localhost user=kantemir password=kantemir dbname=photon port=5432 sslmode=disable"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	fmt.Println("successfully connected to database")

	database.DBConn.AutoMigrate(&post.Post{})
}

func main() {
	// setup settings from env vars and default values
	var s Settings
	err := envconfig.Process("photonserver", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	// create db connection
	initDB()
	//defer database.DBConn.Close()
	// setup fiber app
	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/posts", func(c *fiber.Ctx) {
		posts := post.GetPosts()
		c.Send(posts)
	})

	/* app.Post("/post", func(c *fiber.Ctx) {
		post := post.CreatePost(c)
		c.Send(post)
	}) */

	app.Post("/post", func(c *fiber.Ctx) {
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
			c.Send(createdPost)
		} else {
			c.Status(500).Send("failed to parse the submitted form")
		}
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5600"
	}
	log.Println(app.Listen(port))
}
