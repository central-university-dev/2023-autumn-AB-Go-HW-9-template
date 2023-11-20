package main

import (
	"context"
	"grpc_workshop/3/service" // Import the generated protobuf code
	"log"
	"net"

	"google.golang.org/grpc"
)

// server is used to implement simple.GreetingService.
type server struct {
	service.UnimplementedGreetingServiceServer
}

// SayHello implements simple.GreetingService.
func (s *server) SayHello(ctx context.Context, in *service.HelloRequest) (*service.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &service.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterGreetingServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
