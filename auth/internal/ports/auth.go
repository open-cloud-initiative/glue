package ports

import (
	"context"
	"time"

	"github.com/open-cloud-initiative/glue/auth/internal/models"

	"github.com/google/uuid"
)

// Auth defines the authentication port interface.
type Auth interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	// GetUser retrieves a user by ID.
	GetUser(ctx context.Context, id uuid.UUID) (models.User, error)
	// GetUserByEmail retrieves a user by email.
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	// UpdateUser updates a user.
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	// DeleteUser deletes a user by ID.
	DeleteUser(ctx context.Context, id uuid.UUID) error
	// LinkAccount links an account to a user.
	LinkAccount(ctx context.Context, accountID, userID uuid.UUID) error
	// UnlinkAccount unlinks an account from a user.
	UnlinkAccount(ctx context.Context, accountID, userID uuid.UUID) error
	// CreateSession creates a new session.
	CreateSession(ctx context.Context, userID uuid.UUID, expires time.Time) (models.Session, error)
	// GetSession retrieves a session by session token.
	GetSession(ctx context.Context, sessionToken string) (models.Session, error)
	// UpdateSession updates a session.
	UpdateSession(ctx context.Context, session models.Session) (models.Session, error)
	// RefreshSession refreshes a session.
	RefreshSession(ctx context.Context, session models.Session) (models.Session, error)
	// DeleteSession deletes a session by session token.
	DeleteSession(ctx context.Context, sessionToken string) error
	// CreateVerificationToken creates a new verification token.
	CreateVerificationToken(ctx context.Context, verficationToken models.VerificationToken) (models.VerificationToken, error)
	// UseVerficationToken uses a verification token.
	UseVerficationToken(ctx context.Context, identifier, token string) (models.VerificationToken, error)
}
