package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"function/api"

	"google.golang.org/grpc"
)

// Server is used to implement the API.
type Server struct {
	api.UnimplementedGreeterServer
	grpcServer *grpc.Server
}

func New() (*Server, error) {
	log.Printf("Creating gRPC server")

	grpcServer := grpc.NewServer()
	server := &Server{
		UnimplementedGreeterServer: api.UnimplementedGreeterServer{},
		grpcServer:                 grpcServer,
	}
	api.RegisterGreeterServer(grpcServer, server)

	return server, nil
}

func (s *Server) Serve() error {
	log.Printf("Serving gRPC protocol")
	const port = "50051"

	listener, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		return fmt.Errorf("listen on port %s: %w", port, err)
	}

	if err := s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("serve gRPC: %w", err)
	}

	return nil
}

func (s *Server) SayHello(_ context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
}
