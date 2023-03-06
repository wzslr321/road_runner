package main

import (
	"fmt"
	pb "github.com/wzslr321/road_runner/server/users/gen/proto"
	"github.com/wzslr321/road_runner/server/users/settings"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var serverSettings *settings.Server

type server struct {
	pb.UnimplementedAuthServer
}

var devModeEnv = os.Getenv("DEV_MODE") // if not present, dev is used

func init() {
	devMode := settings.Dev
	if devModeEnv == "PROD" {
		devMode = settings.Prod
	}

	serverSettings = settings.GetServerConf(devMode)
}

func main() {

	listen, err := net.Listen("tcp", fmt.Sprintf("%s", serverSettings.Address))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	log.Printf("Server listening on port %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
