package db

import (
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
