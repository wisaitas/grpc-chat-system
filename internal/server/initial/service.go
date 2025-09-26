package initial

import (
	serverService "github.com/wisaitas/grpc-chat-system/internal/server/service"
	db "github.com/wisaitas/grpc-chat-system/internal/server/sqlc"
)

type service struct {
	UserService *serverService.UserService
}

func newService(
	queries *db.Queries,
) *service {
	return &service{
		UserService: serverService.NewUserService(queries),
	}
}
