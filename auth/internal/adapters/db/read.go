package db

import (
	"context"

	"github.com/open-cloud-initiative/glue/auth/internal/models"
	"github.com/open-cloud-initiative/glue/auth/internal/ports"

	"github.com/katallaxie/pkg/dbx"
	"gorm.io/gorm"
)

var _ ports.ReadTx = (*readTxImpl)(nil)

type readTxImpl struct {
	conn *gorm.DB
}

// NewReadTx ...
func NewReadTx() dbx.ReadTxFactory[ports.ReadTx] {
	return func(db *gorm.DB) (ports.ReadTx, error) {
		return &readTxImpl{conn: db}, nil
	}
}

// GetUser retrieves a user by ID.
func (r *readTxImpl) GetUser(ctx context.Context, user *models.User) error {
	return r.conn.WithContext(ctx).First(user, "id = ?", user.Id).Error
}

// GetAccount retrieves an external account by ID.
func (r *readTxImpl) GetAccount(ctx context.Context, account *models.Account) error {
	return r.conn.WithContext(ctx).First(account, "id = ?", account.ID).Error
}
