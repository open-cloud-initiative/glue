package models

import "time"

// Identity
type Identity struct {
	// Id of the identity.
	Id string
	// User id of the identity.
	UserId string
	// Identity data.
	Data map[string]string
	// Provider.
	Provider string
	// Last signed in at.
	LastSignedInAt time.Time
	// Created at.
	CreatedAt time.Time
	// Updated at.
	UpdatedAt time.Time
	// Deleted at.
	DeletedAt time.Time
	// Email.
	Email string
}
