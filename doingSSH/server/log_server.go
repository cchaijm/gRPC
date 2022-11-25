package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func LogServer() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("Received a request: Server Details\n")

		headers, ok := metadata.FromIncomingContext(ctx)

		if ok {
			log.Printf("Received headers: %v\n", headers)
		}

		return handler(ctx, req)
	}
}
