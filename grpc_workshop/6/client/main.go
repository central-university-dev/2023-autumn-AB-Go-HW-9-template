package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpc_workshop/6/service"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := streaming.NewStreamingServiceClient(conn)

	// Stream request
	stream, err := c.StreamMessages(context.Background(), &streaming.StreamRequest{Count: 5})
	if err != nil {
		log.Fatalf("Error on stream messages: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// End of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving message: %v", err)
		}

		log.Printf("Received message: %s", msg.GetMessage())
	}
}
