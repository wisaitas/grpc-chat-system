package service

import (
	db "github.com/wisaitas/grpc-chat-system/internal/server/sqlc"
)

type UserService interface {
}

type userService struct {
	Queries *db.Queries
}

func NewUserService(Queries *db.Queries) *userService {
	return &userService{
		Queries: Queries,
	}
}
