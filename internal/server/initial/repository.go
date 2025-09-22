package initial

import (
	serverRepository "github.com/wisaitas/grpc-chat-system/internal/server/repository"
	"github.com/wisaitas/grpc-chat-system/internal/server/service"
)

type repository struct {
	userRepository serverRepository.Repository
}

type services struct {
	userService *service.UserService
}

func newRepository(
	cfg *config,
) *repository {
	return &repository{
		userRepository: serverRepository.NewRepository(cfg.Postgres),
	}
}

func newServices(repo *repository) *services {
	return &services{
		userService: service.NewUserService(repo.userRepository),
	}
}
