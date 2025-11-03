package db

import (
	"context"

	"github.com/open-cloud-initiative/glue/auth/internal/models"
	"github.com/open-cloud-initiative/glue/auth/internal/ports"

	"github.com/katallaxie/pkg/dbx"
	"gorm.io/gorm"
)

var _ ports.WriteTx = (*writeTxImpl)(nil)

type writeTxImpl struct {
	conn *gorm.DB
}

// NewWriteTx ...
func NewWriteTx() dbx.ReadWriteTxFactory[ports.WriteTx] {
	return func(db *gorm.DB) (ports.WriteTx, error) {
		return &writeTxImpl{conn: db}, nil
	}
}

// CreateUser creates a new user.
func (w *writeTxImpl) CreateUser(ctx context.Context, user *models.User) error {
	return w.conn.WithContext(ctx).Create(user).Error
}

// UpdateUser updates an existing user.
func (w *writeTxImpl) UpdateUser(ctx context.Context, user *models.User) error {
	return w.conn.WithContext(ctx).Save(user).Error
}

// DeleteUser deletes a user by ID.
func (w *writeTxImpl) DeleteUser(ctx context.Context, user *models.User) error {
	return w.conn.WithContext(ctx).Delete(user, "id = ?", user.ID).Error
}

// LinkAccount links an external account to a user.
func (w *writeTxImpl) LinkAccount(ctx context.Context, account *models.Account, user *models.User) error {
	account.UserID = &user.ID
	return w.conn.WithContext(ctx).Save(account).Error
}

// UnlinkAccount unlinks an external account from a user.
func (w *writeTxImpl) UnlinkAccount(ctx context.Context, account *models.Account, user *models.User) error {
	if account.UserID == nil || *account.UserID != user.ID {
		return gorm.ErrRecordNotFound
	}
	account.UserID = nil
	return w.conn.WithContext(ctx).Save(account).Error
}

// CreateAccount creates a new external account.
func (w *writeTxImpl) CreateAccount(ctx context.Context, account *models.Account) error {
	return w.conn.WithContext(ctx).Create(account).Error
}

// UpdateAccount updates an existing external account.
func (w *writeTxImpl) UpdateAccount(ctx context.Context, account *models.Account) error {
	return w.conn.WithContext(ctx).Save(account).Error
}

// DeleteAccount deletes an external account by ID.
func (w *writeTxImpl) DeleteAccount(ctx context.Context, account *models.Account) error {
	return w.conn.WithContext(ctx).Delete(account, "id = ?", account.ID).Error
}
