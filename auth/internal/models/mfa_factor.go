package models

import (
	"time"
)

// MFAFactorStatus is an enum to represent the current state.
type MFAFactorStatus int32

const (
	// Unspecified is an unspecified state.
	MFAFactorStatus_MFA_FACTOR_STATUS_UNSPECIFIED MFAFactorStatus = 0
	// The factor is in the process of being verified.
	MFAFactorStatus_MFA_FACTOR_STATUS_PENDING MFAFactorStatus = 1
	// The factor has been verified.
	MFAFactorStatus_MFA_FACTOR_STATUS_VERIFIED MFAFactorStatus = 2
	// The factor has been rejected.
	MFAFactorStatus_MFA_FACTOR_STATUS_UNVERIFIED MFAFactorStatus = 3
)

// Multi Factor
type MFAFactor struct {
	// Id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Status.
	Status MFAFactorStatus `protobuf:"varint,2,opt,name=status,proto3,enum=oci.cloud.glue.v1.auth.MFAFactorStatus" json:"status,omitempty"`
	// Type.
	Type MFAFactorType `protobuf:"varint,3,opt,name=type,proto3,enum=oci.cloud.glue.v1.auth.MFAFactorType" json:"type,omitempty"`
	// Friendly name.
	FriendlyName string `protobuf:"bytes,4,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	// Web Authn credentials.
	WebAuthnCredential []byte `protobuf:"bytes,5,opt,name=web_authn_credential,json=webAuthnCredential,proto3" json:"web_authn_credential,omitempty"`
	// Phone number.
	PhoneNumber string `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// Last challenged at.
	LastChallengedAt time.Time `protobuf:"bytes,7,opt,name=last_challenged_at,json=lastChallengedAt,proto3" json:"last_challenged_at,omitempty"`
	// Created at.
	CreatedAt time.Time `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Updated at.
	UpdatedAt time.Time `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Deleted at.
	DeletedAt time.Time `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}
