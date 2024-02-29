package repo

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const (
	generalQueryTimeout = 600 * time.Second
)

type JobRepo struct {
	db *gorm.DB
}

func NewJobRepo(db *gorm.DB) JobRepo {
	return JobRepo{db: db}
}
func (r JobRepo) CreateTx(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.db.WithContext(ctx), cancel
}
