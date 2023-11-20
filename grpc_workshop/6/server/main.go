package main

import (
	"fmt"
	"grpc_workshop/6/service"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	streaming.UnimplementedStreamingServiceServer
}

func (s *server) StreamMessages(req *streaming.StreamRequest, stream streaming.StreamingService_StreamMessagesServer) error {
	for i := 0; i < int(req.Count); i++ {
		message := fmt.Sprintf("Message %d", i+1)
		if err := stream.Send(&streaming.StreamResponse{Message: message}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second) // Simulating delay
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	streaming.RegisterStreamingServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
