package initial

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/wisaitas/grpc-chat-system/internal/server"
	middlewareConfig "github.com/wisaitas/grpc-chat-system/internal/server/middleware/config"
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
	config := newConfig()

	services := newService(config)

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
		Config:     config,
	}
}

func (s *Server) GracefulStop() {
	s.Config.Postgres.Close()
	s.Config.Cassandra.Close()

	s.GrpcServer.GracefulStop()

	log.Println("gRPC server stopped")
}
