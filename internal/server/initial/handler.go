package initial

import (
	serverHandler "github.com/wisaitas/grpc-chat-system/internal/server/handler"
	"google.golang.org/grpc"
)

type handler struct {
	authHandler serverHandler.AuthHandler
}

func newHandler(
	grpcServer *grpc.Server,
	services *service,
) {
	handler := &handler{
		authHandler: serverHandler.NewAuthHandler(grpcServer, services.AuthService),
	}

	handler.register()
}

func (h *handler) register() {
	h.authHandler.Register()
	h.authHandler.Login()
}
