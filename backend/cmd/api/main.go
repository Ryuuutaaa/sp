package main

import (
	"log"

	"sp-backend/internal/config"
	"sp-backend/internal/handler"
	"sp-backend/internal/middleware"
	"sp-backend/internal/repository"
	"sp-backend/internal/routers"
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
	settingRepo := repository.NewSettingRepo(gqlClient)
	userRepo := repository.NewUserRepo(gqlClient)

	// Services
	memberService := service.NewMemberService(memberRepo)
	savingsService := service.NewSavingsService(savingsRepo)
	loanService := service.NewLoanService(loanRepo)
	installmentService := service.NewInstallmentService(installmentRepo)
	cashService := service.NewCashService(cashRepo)
	shuService := service.NewShuService(shuRepo)
	settingService := service.NewSettingService(settingRepo)
	userService := service.NewUserService(userRepo)

	// Handlers (logic only, no route mapping)
	handlers := &handler.Handlers{
		Member:      handler.NewMemberHandler(memberService),
		Savings:     handler.NewSavingsHandler(savingsService),
		Loan:        handler.NewLoanHandler(loanService),
		Installment: handler.NewInstallmentHandler(installmentService),
		Cash:        handler.NewCashHandler(cashService),
		Shu:         handler.NewShuHandler(shuService),
		Setting:     handler.NewSettingHandler(settingService),
		User:        handler.NewUserHandler(userService),
	}

	// Router (semua endpoint + JWT + RBAC di 1 tempat)
	rt := routers.NewRouter(
		middleware.RequireJWTAuth(cfg.JWTSecret),
		middleware.RequireRole,
	)

	api := app.Group("/api")
	rt.RegisterAll(api, handlers)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
