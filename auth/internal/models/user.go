package models

import (
	"time"

	"github.com/google/uuid"
)

// MFAFactorType is an enum of the allowed types.
type MFAFactorType int32

const (
	// Unspecified is an unspecified type.
	MFAFactorType_FACTOR_TYPE_UNSPECIFIED MFAFactorType = 0
	// The factor is a time-based one-time password.
	MFAFactorType_FACTOR_TYPE_TOTP MFAFactorType = 1
	// The factor is a biometric.
	MFAFactorType_FACTOR_TYPE_BIOMETRIC MFAFactorType = 2
	// The factor is a hardware token.
	MFAFactorType_FACTOR_TYPE_HARDWARE_TOKEN MFAFactorType = 3
)

// User represents a user in the system.
type User struct {
	// Id of the user.
	Id uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
	// Role is the role of the user.
	Role string `json:"role" gorm:"type:string"`
	// Email is the email of the user.
	Email string `json:"email" gorm:"type:string"`
	// Email verified at.
	EmailVerifiedAt time.Time
	// Phone number of the user.
	PhoneNumber string
	// Phone number verified at.
	PhoneNumberVerifiedAt time.Time
	// Confirmation at.
	ConfirmedAt time.Time
	// Confirmation send at.
	ConfirmationSentAt time.Time
	// Reauthentication at.
	ReauthenticatedAt time.Time
	// Last signed in at.
	LastSignedInAt time.Time `protobuf:"bytes,10,opt,name=last_signed_in_at,json=lastSignedInAt,proto3" json:"last_signed_in_at,omitempty"`
	// App metadata.
	AppMetadata map[string]string `protobuf:"bytes,11,rep,name=app_metadata,json=appMetadata,proto3" json:"app_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User metadata.
	UserMetadata map[string]string `protobuf:"bytes,12,rep,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Banned until.
	BannedUntil time.Time `protobuf:"bytes,13,opt,name=banned_until,json=bannedUntil,proto3" json:"banned_until,omitempty"`
	// Created at.
	CreatedAt time.Time `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Updated at.
	UpdatedAt time.Time `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Deleted at.
	DeletedAt time.Time `protobuf:"bytes,16,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// Is anonymous.
	IsAnonymous bool `protobuf:"varint,17,opt,name=is_anonymous,json=isAnonymous,proto3" json:"is_anonymous,omitempty"`
	// Identities.
	Identities []*Identity `protobuf:"bytes,18,rep,name=identities,proto3" json:"identities,omitempty"`
	// MFA factors.
	MfaFactors []*MFAFactor
}
