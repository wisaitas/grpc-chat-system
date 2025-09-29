package handler

import (
	authPb "github.com/wisaitas/grpc-chat-system/internal/server/protogen/auth"
	"google.golang.org/grpc"
)

type AuthHandler interface {
	Login()
	Register()
}

type authHandler struct {
	grpcServer  *grpc.Server
	authService authPb.AuthServiceServer
}

func NewAuthHandler(
	grpcServer *grpc.Server,
	authService authPb.AuthServiceServer,
) *authHandler {
	return &authHandler{
		grpcServer:  grpcServer,
		authService: authService,
	}
}

func (h *authHandler) Register() {
	authPb.RegisterAuthServiceServer(h.grpcServer, h.authService)
}

func (h *authHandler) Login() {
	authPb.RegisterAuthServiceServer(h.grpcServer, h.authService)
}
