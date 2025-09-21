package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryLoggingInterceptor logs all unary RPC calls
func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// start := time.Now()

		// Call the handler
		resp, err := handler(ctx, req)

		// Calculate duration
		// duration := time.Since(start)

		// Log the result
		if err != nil {
			st, _ := status.FromError(err)
			log.Println("gRPC request failed", st.Code().String(), err.Error())
		} else {
			log.Println("gRPC request completed")
		}

		return resp, err
	}
}

// StreamLoggingInterceptor logs all streaming RPC calls
func StreamLoggingInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// start := time.Now()

		log.Println("gRPC stream started")

		// Call the handler
		err := handler(srv, stream)

		// Calculate duration
		// duration := time.Since(start)

		// Log the result
		if err != nil {
			st, _ := status.FromError(err)
			log.Println("gRPC stream failed", st.Code().String(), err.Error())
		} else {
			log.Println("gRPC stream completed")
		}

		return err
	}
}
