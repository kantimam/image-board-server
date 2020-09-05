package main

import (
	"fmt"
	"log"
	"os"
	"server/database"

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
	// setup fiber app
	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/getepisodes", func(c *fiber.Ctx) {

	})

	app.Get("/search", func(c *fiber.Ctx) {
		c.Status(201).JSON(fiber.Map{
			"test":  "passed",
			"hello": "world",
		})
	})

	app.Get("/getstream", func(c *fiber.Ctx) {
		videoID := c.Query("id")

		if videoID == "" {
			c.Send(`please provide a valid id`)
			return
		}
		c.Send("success")
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5600"
	}
	log.Println(app.Listen(port))
}
