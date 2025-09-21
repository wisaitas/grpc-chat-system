package repository

import (
	"context"

	"github.com/wisaitas/grpc-chat-system/internal/server/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type repository struct {
	postgres *gorm.DB
}

func NewRepository(
	postgres *gorm.DB,
) Repository {
	return &repository{
		postgres: postgres,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *model.User) error {
	return r.postgres.WithContext(ctx).Create(user).Error
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.postgres.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}
