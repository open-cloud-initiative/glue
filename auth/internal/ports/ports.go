package ports

import (
	"context"

	"github.com/open-cloud-initiative/glue/auth/internal/models"
)

// ReadTx is the interface for read-only transactions.
type ReadTx interface {
	// GetUser retrieves a user by ID.
	GetUser(ctx context.Context, user *models.User) error
	// GetAccount retrieves an external account by ID.
	GetAccount(ctx context.Context, account *models.Account) error
}

// WriteTx is the interface for read-write transactions.
type WriteTx interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user *models.User) error
	// UpdateUser updates an existing user.
	UpdateUser(ctx context.Context, user *models.User) error
	// DeleteUser deletes a user by ID.
	DeleteUser(ctx context.Context, user *models.User) error
	// LinkAccount links an external account to a user.
	LinkAccount(ctx context.Context, account *models.Account, user *models.User) error
	// UnlinkAccount unlinks an external account from a user.
	UnlinkAccount(ctx context.Context, account *models.Account, user *models.User) error
	// CreateAccount creates a new external account.
	CreateAccount(ctx context.Context, account *models.Account) error
	// UpdateAccount updates an existing external account.
	UpdateAccount(ctx context.Context, account *models.Account) error
	// DeleteAccount deletes an external account by ID.
	DeleteAccount(ctx context.Context, account *models.Account) error
}
