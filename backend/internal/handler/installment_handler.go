package handler

import (
	"sp-backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type InstallmentHandler struct {
	service domain.InstallmentService
}

func NewInstallmentHandler(service domain.InstallmentService) *InstallmentHandler {
	return &InstallmentHandler{service: service}
}

func (h *InstallmentHandler) Register(router fiber.Router) {
	installments := router.Group("/installments")
	installments.Get("/loan/:loanId", h.GetByLoanID)
	installments.Post("/:id/pay", h.Pay)
}

func (h *InstallmentHandler) GetByLoanID(c *fiber.Ctx) error {
	loanID := c.Params("loanId")
	installments, err := h.service.GetByLoanID(loanID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(installments)
}

func (h *InstallmentHandler) Pay(c *fiber.Ctx) error {
	id := c.Params("id")
	var input domain.PayInstallmentInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	installment, err := h.service.Pay(id, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(installment)
}
