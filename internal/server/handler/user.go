package handler

import (
	userPb "github.com/wisaitas/grpc-chat-system/internal/server/protogen/user"
	"github.com/wisaitas/grpc-chat-system/internal/server/service"
	"google.golang.org/grpc"
)

type UserHandler interface {
	Register()
}

type userHandler struct {
	grpcServer  *grpc.Server
	userService *service.UserService
}

func NewUserHandler(
	grpcServer *grpc.Server,
	userService *service.UserService,
) *userHandler {
	return &userHandler{
		grpcServer:  grpcServer,
		userService: userService,
	}
}

func (h *userHandler) Register() {
	userPb.RegisterUserServiceServer(h.grpcServer, h.userService)
}
