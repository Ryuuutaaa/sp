package handler

import (
	"sp-backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	service domain.MemberService
}

func NewMemberHandler(service domain.MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) Register(router fiber.Router) {
	members := router.Group("/members")
	members.Get("/", h.GetAll)
	members.Get("/:id", h.GetByID)
	members.Post("/", h.Create)
	members.Patch("/:id/verify", h.Verify)
}

func (h *MemberHandler) GetAll(c *fiber.Ctx) error {
	members, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(members)
}

func (h *MemberHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	member, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (h *MemberHandler) Create(c *fiber.Ctx) error {
	var input domain.CreateMemberInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	member, err := h.service.Register(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(member)
}

func (h *MemberHandler) Verify(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Status     string `json:"status"`
		VerifiedBy string `json:"verifiedBy"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	member, err := h.service.Verify(id, body.Status, body.VerifiedBy)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}
