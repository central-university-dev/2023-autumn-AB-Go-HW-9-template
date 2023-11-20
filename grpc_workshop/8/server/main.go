package main

import (
	"context"
	"fmt"
	"grpc_workshop/8/interceptor"
	validation "grpc_workshop/8/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	validation.UnimplementedValidationServiceServer
}

func (s *server) CreateUser(ctx context.Context, req *validation.CreateUserRequest) (*validation.CreateUserResponse, error) {
	return &validation.CreateUserResponse{
		Message: fmt.Sprintf("User %s created successfully!", req.GetName()),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.ValidateInterceptor))
	validation.RegisterValidationServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
