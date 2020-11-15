package main

import (
	"fmt"
	"log"
	"os"
	"server/database"
	"server/models/post"
	"server/models/user"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

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

	database.DBConn.AutoMigrate(&post.Post{}, &user.User{})
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

	// create fiber app
	app := fiber.New()
	// setupt middlewares
	app.Static("/static", "./static")
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(logger.New())
	// setup routes
	routes.CreateRoutes(app)

	// start listening for requests
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5600"
	}
	log.Println(app.Listen(fmt.Sprintf(":%s", port)))
}
