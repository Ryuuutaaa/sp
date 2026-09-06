package handler

import (
	"sp-backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type SettingHandler struct {
	service domain.SettingService
}

func NewSettingHandler(service domain.SettingService) *SettingHandler {
	return &SettingHandler{service: service}
}

func (h *SettingHandler) GetAll(c *fiber.Ctx) error {
	settings, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(settings)
}

func (h *SettingHandler) Upsert(c *fiber.Ctx) error {
	key := c.Params("key")
	var body struct {
		Value       string  `json:"value"`
		Description *string `json:"description"`
	}
	if err := c.BodyParser(&body); err != nil || key == "" {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	setting, err := h.service.Upsert(domain.UpsertSettingInput{Key: key, Value: body.Value, Description: body.Description})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(setting)
}
