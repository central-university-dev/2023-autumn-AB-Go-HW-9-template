package interceptor

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor measures the duration of RPC calls
func UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)
	return h, err
}

// UnaryClientInterceptor measures the duration of RPC calls
func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		method,
		time.Since(start),
		err)
	return err
}

// UnaryMetadataClientInterceptor for adding metadata
func UnaryMetadataClientInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Create a new context with the ID metadata
	md := metadata.Pairs("id", strconv.Itoa(rand.Intn(100)))
	newCtx := metadata.NewOutgoingContext(ctx, md)

	return invoker(newCtx, method, req, reply, cc, opts...)
}

// UnaryMetadataServerInterceptor for reading metadata
func UnaryMetadataServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// Log the ID from metadata, if present
		if correlationIDs := md["id"]; len(correlationIDs) > 0 {
			log.Printf("ID: %s\n", correlationIDs[0])
		}
	}
	return handler(ctx, req)
}
