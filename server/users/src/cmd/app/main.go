package main

import (
	"fmt"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	service := NewUserService()
	service = NewLoggingService(service)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUsersServer(server, newGrpcServer())
	server.Serve(listener)
}
