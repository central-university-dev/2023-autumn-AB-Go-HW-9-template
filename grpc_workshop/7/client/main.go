package main

import (
	"context"
	"grpc_workshop/7/bidirectional"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := bidirectional.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// Channel to handle incoming messages
	done := make(chan bool)

	// Goroutine to handle incoming messages
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// End of stream
				done <- true
				return
			}
			if err != nil {
				log.Printf("Failed to receive a message : %v", err)
				done <- true
				return
			}
			log.Printf("Received message: %s", in.GetText())
		}
	}()

	// Sending messages in a loop
	for i := 0; i < 5; i++ {
		message := &bidirectional.ChatMessage{Text: "Hello, Server! Message " + strconv.Itoa(i+1)}
		if err := stream.Send(message); err != nil {
			log.Printf("Failed to send a message: %v", err)
			break
		}
		time.Sleep(1 * time.Second)
	}

	// Close the send direction of the stream
	stream.CloseSend()

	// Wait for the receiver goroutine to finish
	<-done
}
