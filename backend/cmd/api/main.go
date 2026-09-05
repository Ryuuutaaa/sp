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
	savingsRepo := repository.NewSavingsRepo(gqlClient)
	loanRepo := repository.NewLoanRepo(gqlClient)
	installmentRepo := repository.NewInstallmentRepo(gqlClient)
	cashRepo := repository.NewCashRepo(gqlClient)
	shuRepo := repository.NewShuRepo(gqlClient)

	// Services
	memberService := service.NewMemberService(memberRepo)
	savingsService := service.NewSavingsService(savingsRepo)
	loanService := service.NewLoanService(loanRepo)
	installmentService := service.NewInstallmentService(installmentRepo)
	cashService := service.NewCashService(cashRepo)
	shuService := service.NewShuService(shuRepo)

	// Handlers
	memberHandler := handler.NewMemberHandler(memberService)
	savingsHandler := handler.NewSavingsHandler(savingsService)
	loanHandler := handler.NewLoanHandler(loanService)
	installmentHandler := handler.NewInstallmentHandler(installmentService)
	cashHandler := handler.NewCashHandler(cashService)
	shuHandler := handler.NewShuHandler(shuService)

	// Routes
	api := app.Group("/api")
	memberHandler.Register(api)
	savingsHandler.Register(api)
	loanHandler.Register(api)
	installmentHandler.Register(api)
	cashHandler.Register(api)
	shuHandler.Register(api)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
