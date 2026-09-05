package handler

import (
	"sp-backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type LoanHandler struct {
	service domain.LoanService
}

func NewLoanHandler(service domain.LoanService) *LoanHandler {
	return &LoanHandler{service: service}
}

func (h *LoanHandler) GetAll(c *fiber.Ctx) error {
	loans, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(loans)
}

func (h *LoanHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	loan, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(loan)
}

func (h *LoanHandler) Apply(c *fiber.Ctx) error {
	var input domain.CreateLoanInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	loan, err := h.service.Apply(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(loan)
}

func (h *LoanHandler) Approve(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		ApprovedBy string `json:"approvedBy"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	loan, err := h.service.Approve(id, body.ApprovedBy)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(loan)
}

func (h *LoanHandler) Reject(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		RejectedBy string `json:"rejectedBy"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	loan, err := h.service.Reject(id, body.RejectedBy)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(loan)
}

func (h *LoanHandler) Disburse(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		DisbursedBy string `json:"disbursedBy"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	loan, err := h.service.Disburse(id, body.DisbursedBy)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(loan)
}
