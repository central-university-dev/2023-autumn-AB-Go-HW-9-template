package main

import (
	"context"
	"grpc_workshop/4/quotes"
	"grpc_workshop/5/interceptor"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	quotes.UnimplementedQuotesServiceServer
}

func (s *server) GetQuotes(ctx context.Context, in *quotes.QuotesRequest) (*quotes.QuotesResponse, error) {
	if in.NumberOfQuotes < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "NumberOfQuotes must be positive")
	}

	// Dummy quotes for demonstration
	dummyQuotes := []string{"Quote 1", "Quote 2", "Quote 3"}
	if int(in.NumberOfQuotes) > len(dummyQuotes) {
		return nil, status.Errorf(codes.OutOfRange, "Requested more quotes than available")
	}

	return &quotes.QuotesResponse{Quotes: dummyQuotes[:in.NumberOfQuotes]}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.UnaryServerInterceptor, interceptor.UnaryMetadataServerInterceptor))
	quotes.RegisterQuotesServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
