package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// VerificationToken is a verification token for a user.
type VerificationToken struct {
	// Token is the unique identifier of the token.
	Token string `json:"token" gorm:"primaryKey"`
	// Identifier is the identifier of the token.
	Identifier string `json:"identifier"`
	// ExpiresAt is the expiry time of the token.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt is the creation time of the token.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the token.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the token.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// CsrfToken is a CSRF token for a session.
type CsrfToken struct {
	// ID is the unique identifier of the CSRF token.
	ID uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
	// Token is the unique identifier of the token.
	Token string `json:"token"`
	// ExpiresAt is the expiry time of the token.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt is the creation time of the token.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the token.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the token.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Session represents a user session.
type Session struct {
	// ID is the unique identifier of the session.
	ID uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
	// SessionToken is the token of the session.
	SessionToken string `json:"session_token"`
	// CsrfToken is the CSRF token of the session.
	CsrfToken CsrfToken `json:"csrf_token"`
	// CsrfTokenID is the CSRF token ID of the session.
	CsrfTokenID uuid.UUID `json:"csrf_token_id"`
	// UserID is the user ID of the session.
	UserID uuid.UUID `json:"user_id"`
	// User is the user of the session.
	User User `json:"user"`
	// ExpiresAt is the expiry time of the session.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt is the creation time of the session.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the session.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the session.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
