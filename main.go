package main

import (
	"log"
	"time"

	"alfredtomo.dev/todos-app/database"
	"alfredtomo.dev/todos-app/database/migration"
	"alfredtomo.dev/todos-app/docs"
	"alfredtomo.dev/todos-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
)

// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	hello@alfredtomo.dev
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "This is a sample todo server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// INITIATE DB
	database.Init()

	// RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// 3 requests per 10 seconds max
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Initiate Route from routes
	routes.RouteInit(app)
	routes.TodoRoutes(app)

	log.Fatal(app.Listen("localhost:8080"))
}
