package handler

import (
	"sp-backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type SavingsHandler struct {
	service domain.SavingsService
}

func NewSavingsHandler(service domain.SavingsService) *SavingsHandler {
	return &SavingsHandler{service: service}
}

func (h *SavingsHandler) Register(router fiber.Router) {
	savings := router.Group("/savings")
	savings.Get("/types", h.GetTypes)
	savings.Get("/transactions", h.GetTransactions)
	savings.Post("/deposit", h.Deposit)
	savings.Post("/withdraw", h.Withdraw)
}

func (h *SavingsHandler) GetTypes(c *fiber.Ctx) error {
	types, err := h.service.GetTypes()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(types)
}

func (h *SavingsHandler) GetTransactions(c *fiber.Ctx) error {
	memberID := c.Query("memberId")
	txs, err := h.service.GetTransactions(memberID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(txs)
}

func (h *SavingsHandler) Deposit(c *fiber.Ctx) error {
	var input domain.CreateSavingsInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	tx, err := h.service.Deposit(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(tx)
}

func (h *SavingsHandler) Withdraw(c *fiber.Ctx) error {
	var input domain.CreateSavingsInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	tx, err := h.service.Withdraw(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(tx)
}
