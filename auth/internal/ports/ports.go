package ports

import (
	"context"

	"github.com/open-cloud-initiative/glue/auth/internal/models"
)

// ReadTx is the interface for read-only transactions.
type ReadTx interface {
	// GetUser retrieves a user by ID.
	GetUser(ctx context.Context, user *models.User) error
}

// WriteTx is the interface for read-write transactions.
type WriteTx interface {
	ReadTx
}
