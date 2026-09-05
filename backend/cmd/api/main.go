package main

import (
	"log"

	"sp-backend/internal/config"
	"sp-backend/internal/handler"
	"sp-backend/internal/repository"
	"sp-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.Load()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// GraphQL Client
	gqlClient := repository.NewGQLClient(cfg.GQLURL)

	// Repositories
	memberRepo := repository.NewMemberRepo(gqlClient)

	// Services
	memberService := service.NewMemberService(memberRepo)

	// Handlers
	memberHandler := handler.NewMemberHandler(memberService)

	// Routes
	api := app.Group("/api")
	memberHandler.Register(api)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
