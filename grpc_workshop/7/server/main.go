package main

import (
	"grpc_workshop/7/bidirectional"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	bidirectional.UnimplementedChatServiceServer
}

func (s *server) Chat(stream bidirectional.ChatService_ChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil // End of stream
		}
		if err != nil {
			log.Fatalf("Failed to receive a message : %v", err)
			return err
		}

		log.Printf("Received message: %s", in.GetText())
		if err := stream.Send(in); err != nil { // Echo the received message back to the client
			log.Fatalf("Failed to send a message: %v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bidirectional.RegisterChatServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
