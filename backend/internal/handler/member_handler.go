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

func (h *MemberHandler) GetAll(c *fiber.Ctx) error {
	members, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(members)
}

func (h *MemberHandler) Me(c *fiber.Ctx) error {
	memberID, _ := c.Locals("memberId").(string)
	if memberID == "" {
		return c.Status(404).JSON(fiber.Map{"error": "member profile not linked"})
	}
	member, err := h.service.GetByID(memberID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if member == nil {
		return c.Status(404).JSON(fiber.Map{"error": "member not found"})
	}
	return c.JSON(member)
}

func (h *MemberHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	member, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (h *MemberHandler) RegisterPublic(c *fiber.Ctx) error {
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

func (h *MemberHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var input domain.UpdateMemberInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	member, err := h.service.Update(id, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
}

func (h *MemberHandler) UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&body); err != nil || (body.Status != "active" && body.Status != "inactive") {
		return c.Status(400).JSON(fiber.Map{"error": "invalid status"})
	}
	member, err := h.service.UpdateStatus(id, body.Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(member)
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
