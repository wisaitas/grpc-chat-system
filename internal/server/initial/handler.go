package initial

import (
	serverHandler "github.com/wisaitas/grpc-chat-system/internal/server/handler"
	"google.golang.org/grpc"
)

type handler struct {
	userHandler serverHandler.UserHandler
}

func newHandler(
	grpcServer *grpc.Server,
	services *services,
) {
	handler := &handler{
		userHandler: serverHandler.NewUserHandler(grpcServer, services.userService),
	}

	handler.Register()
}

func (h *handler) Register() {
	h.userHandler.Register()
}
