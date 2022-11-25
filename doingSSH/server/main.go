package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	"github.com/cchaijm/gRPC/doingSSH/proto"
)

var addr string = "0.0.0.0:50050"

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	defer listen.Close()

	log.Printf("Listening on %s\n", addr)

	opts := []grpc.ServerOption{}

	tls := false // change that to true if needed
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	opts = append(opts, grpc.ChainUnaryInterceptor(LogServer(), LogHeader()))

	s := grpc.NewServer(opts...)
	proto.RegisterDoingSSHServer(s, &Server{})

	defer s.Stop()
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
