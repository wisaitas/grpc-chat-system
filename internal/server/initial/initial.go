package initial

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/wisaitas/grpc-chat-system/internal/server"
	middlewareConfig "github.com/wisaitas/grpc-chat-system/internal/server/middleware/config"
	db "github.com/wisaitas/grpc-chat-system/internal/server/sqlc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	if err := env.Parse(&server.Config); err != nil {
		log.Fatalf("failed to parse server config: %v", err)
	}
}

type Server struct {
	GrpcServer *grpc.Server
	Config     *config
}

func New() *Server {
	cfg := newConfig()

	queries := db.New(cfg.DB)
	services := newService(queries)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middlewareConfig.UnaryRecoveryInterceptor(),
			middlewareConfig.UnaryLoggingInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			middlewareConfig.StreamRecoveryInterceptor(),
			middlewareConfig.StreamLoggingInterceptor(),
		),
	)
	reflection.Register(grpcServer)

	newHandler(grpcServer, services)

	return &Server{
		GrpcServer: grpcServer,
		Config:     cfg,
	}
}

func (s *Server) GracefulStop() {
	s.Config.DB.Close()

	s.GrpcServer.GracefulStop()

	log.Println("gRPC server stopped")
}
