package initial

import (
	authPb "github.com/wisaitas/grpc-chat-system/internal/server/protogen/auth"
	serverService "github.com/wisaitas/grpc-chat-system/internal/server/service"
	db "github.com/wisaitas/grpc-chat-system/internal/server/sqlc"
)

type service struct {
	AuthService authPb.AuthServiceServer
}

func newService(
	config *config,
) *service {
	queries := db.New(config.Postgres)
	return &service{
		AuthService: serverService.NewAuthService(queries),
	}
}
