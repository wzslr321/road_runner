package main

import (
	"github.com/wzslr321/road_runner/server/users/src/api"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	service := api.NewUsersService()
	service = api.NewLoggingService(service)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUsersServer(server, api.NewServer(service))
	server.Serve(listener)
}
