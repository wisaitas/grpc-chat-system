package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wisaitas/grpc-chat-system/internal/server"
	serverInitial "github.com/wisaitas/grpc-chat-system/internal/server/initial"
)

func main() {
	serverInitial := serverInitial.New()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.Config.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", server.Config.Server.Port, err)
	}

	log.Printf("gRPC server starting on port %s", server.Config.Server.Port)
	if err := serverInitial.GrpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
