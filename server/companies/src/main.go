package main

import (
	"github.com/wzslr321/road_runner/server/companies/src/api"
	pb "github.com/wzslr321/road_runner/server/companies/src/proto-gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
)

func main() {
	service := api.NewCompaniesService()
	service = api.NewLoggingService(service)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{}),
	)
	pb.RegisterCompaniesServer(server, api.NewServer(service))
	server.Serve(listener)
}
