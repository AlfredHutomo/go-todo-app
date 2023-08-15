package main

import (
	"log"

	"alfredtomo.dev/todos-app/database"
	"alfredtomo.dev/todos-app/database/migration"
	"alfredtomo.dev/todos-app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    // INITIATE DB
    database.Init()

    // RUN MIGRATION
    migration.RunMigration()

    app := fiber.New()

    // Initiate Route from routes
    routes.RouteInit(app)
    routes.TodoRoutes(app)

    log.Fatal(app.Listen("localhost:8080"))
}