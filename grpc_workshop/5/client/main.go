package main

import (
	"context"
	"grpc_workshop/4/quotes"
	"grpc_workshop/5/interceptor"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(
		insecure.NewCredentials()), grpc.WithBlock(),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryClientInterceptor, interceptor.UnaryMetadataClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := quotes.NewQuotesServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetQuotes(ctx, &quotes.QuotesRequest{NumberOfQuotes: 10})
	if err != nil {
		// Handle specific gRPC error
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.InvalidArgument:
				log.Printf("Invalid argument: %v", e.Message())
			case codes.OutOfRange:
				log.Printf("Out of range: %v", e.Message())
			default:
				log.Printf("Unknown error: %v", e.Message())
			}
		} else {
			log.Fatalf("Failed to get quotes: %v", err)
		}
		return
	}
	log.Printf("Quotes: %v", r.GetQuotes())
}
