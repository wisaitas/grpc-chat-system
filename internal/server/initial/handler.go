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
	services *service,
) {
	handler := &handler{
		userHandler: serverHandler.NewUserHandler(grpcServer, services.UserService),
	}

	handler.Register()
}

func (h *handler) Register() {
	h.userHandler.Register()
}
