package controllers

import (
	"github.com/open-cloud-initiative/glue/auth/internal/ports"

	"github.com/gofiber/fiber/v3"
	"github.com/katallaxie/pkg/dbx"
)

// UserController handles user-related operations.
type UserController struct {
	store dbx.Database[ports.ReadTx, ports.WriteTx]
}

// NewUserController creates a new UserController.
func NewUserController(store dbx.Database[ports.ReadTx, ports.WriteTx]) *UserController {
	return &UserController{store: store}
}

// GetUser retrieves a user by ID.
func (uc *UserController) GetUser(ctx fiber.Ctx) {
}
