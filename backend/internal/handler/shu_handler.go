package handler

import (
	"strconv"

	"sp-backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type ShuHandler struct {
	service domain.ShuService
}

func NewShuHandler(service domain.ShuService) *ShuHandler {
	return &ShuHandler{service: service}
}

func (h *ShuHandler) GetByYear(c *fiber.Ctx) error {
	yearStr := c.Params("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid year"})
	}
	shus, err := h.service.GetByYear(year)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	// Anggota hanya boleh lihat SHU miliknya sendiri.
	if role, _ := c.Locals("userRole").(string); role == "anggota" {
		memberID, _ := c.Locals("memberId").(string)
		filtered := make([]domain.ShuDistribution, 0)
		for _, s := range shus {
			if s.MemberID == memberID {
				filtered = append(filtered, s)
			}
		}
		return c.JSON(filtered)
	}
	return c.JSON(shus)
}

func (h *ShuHandler) Calculate(c *fiber.Ctx) error {
	yearStr := c.Query("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid year"})
	}
	shus, err := h.service.Calculate(year)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(shus)
}
