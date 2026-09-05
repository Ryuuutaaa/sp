package handler

import (
	"sp-backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type CashHandler struct {
	service domain.CashService
}

func NewCashHandler(service domain.CashService) *CashHandler {
	return &CashHandler{service: service}
}

func (h *CashHandler) GetAll(c *fiber.Ctx) error {
	txs, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(txs)
}

func (h *CashHandler) Record(c *fiber.Ctx) error {
	var input domain.CreateCashInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	tx, err := h.service.Record(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(tx)
}
