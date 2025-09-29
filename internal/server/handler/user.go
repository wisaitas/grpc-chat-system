package handler

import (
	"google.golang.org/grpc"
)

type UserHandler interface {
	Login()
	Register()
}

type userHandler struct {
	grpcServer *grpc.Server
	// userService service.AuthService
}

func NewUserHandler(
	grpcServer *grpc.Server,
	// userService service.AuthService,
) *userHandler {
	return &userHandler{
		grpcServer: grpcServer,
		// userService: userService,
	}
}
