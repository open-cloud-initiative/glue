package models

import (
	"encoding/gob"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	gob.Register(&Account{})
	gob.Register(&User{})
}

// AccountType represents the type of an account.
type AccountType string

// ErrUnimplemented is returned when a method is not implemented.
var ErrUnimplemented = errors.New("not implemented")

const (
	// AccountTypeOAuth2 represents an OAuth2 account type.
	AccountTypeOAuth2 AccountType = "oauth2"
	// AccountTypeOIDC represents an OIDC account type.
	AccountTypeOIDC AccountType = "oidc"
	// AccountTypeSAML represents a SAML account type.
	AccountTypeSAML AccountType = "saml"
	// AccountTypeEmail represents an email account type.
	AccountTypeEmail AccountType = "email"
	// AccountTypeWebAuthn represents a WebAuthn account type.
	AccountTypeWebAuthn AccountType = "webauthn"
)

// Account represents an external account linked to a user.
type Account struct {
	// ID is the unique identifier of the account.
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;column:id;default:gen_random_uuid();"`
	// Type is the type of the account.
	Type AccountType `json:"type" validate:"required"`
	// Provider is the provider of the account.
	Provider string `json:"provider" validate:"required"`
	// ProviderAccountID is the account ID in the provider.
	ProviderAccountID *string `json:"provider_account_id"`
	// RefreshToken is the refresh token of the account.
	RefreshToken *string `json:"refresh_token"`
	// AccessToken is the access token of the account.
	AccessToken *string `json:"access_token"`
	// ExpiresAt is the expiry time of the account.
	ExpiresAt *time.Time `json:"expires_at"`
	// TokenType is the token type of the account.
	TokenType *string `json:"token_type"`
	// Scope is the scope of the account.
	Scope *string `json:"scope"`
	// IDToken is the ID token of the account.
	IDToken *string `json:"id_token"`
	// SessionState is the session state of the account.
	SessionState string `json:"session_state"`
	// UserID is the user ID of the account.
	UserID *uuid.UUID `json:"user_id"`
	//  User is the user of the account.
	User User `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// CreatedAt is the creation time of the account.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the account.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the account.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
