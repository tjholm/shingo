package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	v1 "github.com/tjholm/shingo/pkg/api/shingo/v1"
	"github.com/tjholm/shingo/pkg/events"
	"google.golang.org/grpc"
)

func main() {
	// start the events server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	v1.RegisterEventsServiceServer(grpcServer, events.New())
	httpServ := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool { return true }))
	fmt.Println("server listenting on :50051")
	http.Serve(lis, httpServ)
}
