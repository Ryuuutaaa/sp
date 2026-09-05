package routers

import (
	"sp-backend/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	RequireJWTAuth fiber.Handler
	RequireRole    func(...string) fiber.Handler
}

func NewRouter(jwtAuth fiber.Handler, requireRole func(...string) fiber.Handler) *Router {
	return &Router{
		RequireJWTAuth: jwtAuth,
		RequireRole:    requireRole,
	}
}

const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
	RoleTeller     = "teller"
	RoleAnggota    = "anggota"
)

// ==========================================
// MEMBER ROUTES
// ==========================================

func (rt *Router) RegisterMemberRoutes(r fiber.Router, h *handler.MemberHandler) {
	members := r.Group("/members")
	members.Post("/register", h.RegisterPublic)
	members.Get("/", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin, RoleTeller), h.GetAll)
	members.Get("/:id", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin, RoleTeller), h.GetByID)
	members.Post("/", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.Create)
	members.Patch("/:id/verify", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.Verify)
}

// ==========================================
// SAVINGS ROUTES (Teller operasional)
// ==========================================

func (rt *Router) RegisterSavingsRoutes(r fiber.Router, h *handler.SavingsHandler) {
	savings := r.Group("/savings")
	savings.Get("/types", rt.RequireJWTAuth, h.GetTypes)
	savings.Get("/transactions", rt.RequireJWTAuth, h.GetTransactions)
	savings.Post("/deposit", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin, RoleTeller), h.Deposit)
	savings.Post("/withdraw", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin, RoleTeller), h.Withdraw)
}

// ==========================================
// LOAN ROUTES
// ==========================================

func (rt *Router) RegisterLoanRoutes(r fiber.Router, h *handler.LoanHandler) {
	loans := r.Group("/loans")
	loans.Get("/", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin, RoleTeller), h.GetAll)
	loans.Get("/:id", rt.RequireJWTAuth, h.GetByID)
	loans.Post("/apply", rt.RequireJWTAuth, rt.RequireRole(RoleAnggota), h.Apply)
	loans.Patch("/:id/approve", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.Approve)
	loans.Patch("/:id/reject", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.Reject)
	loans.Patch("/:id/disburse", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.Disburse)
}

// ==========================================
// INSTALLMENT ROUTES
// ==========================================

func (rt *Router) RegisterInstallmentRoutes(r fiber.Router, h *handler.InstallmentHandler) {
	installments := r.Group("/installments")
	installments.Get("/loan/:loanId", rt.RequireJWTAuth, h.GetByLoanID)
	installments.Post("/:id/pay", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin, RoleTeller), h.Pay)
}

// ==========================================
// CASH ROUTES (Admin only)
// ==========================================

func (rt *Router) RegisterCashRoutes(r fiber.Router, h *handler.CashHandler) {
	cash := r.Group("/cash")
	cash.Get("/", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.GetAll)
	cash.Post("/", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.Record)
}

// ==========================================
// SHU ROUTES (Admin only)
// ==========================================

func (rt *Router) RegisterShuRoutes(r fiber.Router, h *handler.ShuHandler) {
	shu := r.Group("/shu")
	shu.Get("/:year", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.GetByYear)
	shu.Post("/calculate", rt.RequireJWTAuth, rt.RequireRole(RoleSuperAdmin, RoleAdmin), h.Calculate)
}

// RegisterAll wires every domain route at once.
func (rt *Router) RegisterAll(api fiber.Router, h *handler.Handlers) {
	rt.RegisterMemberRoutes(api, h.Member)
	rt.RegisterSavingsRoutes(api, h.Savings)
	rt.RegisterLoanRoutes(api, h.Loan)
	rt.RegisterInstallmentRoutes(api, h.Installment)
	rt.RegisterCashRoutes(api, h.Cash)
	rt.RegisterShuRoutes(api, h.Shu)
}
