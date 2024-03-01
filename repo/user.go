package repo

import (
	"context"
	"go-rest-api-with-test/model"
	"gorm.io/gorm"
	"time"
)

const (
	generalQueryTimeout = 600 * time.Second
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db: db}
}
func (r UserRepo) CreateTx(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func (r UserRepo) Create(ctx context.Context, ob *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()
	err := r.db.WithContext(ctx).Model(&model.User{}).Create(ob).Error
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepo) GetOne(ctx context.Context, id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	defer cancel()
	var (
		result *model.User
		err    error
	)

	err = r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
