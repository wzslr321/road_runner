package main

import (
	"github.com/gocql/gocql"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"github.com/wzslr321/road_runner/server/users/src/settings"
	"go.uber.org/zap"
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

	listen, err := net.Listen("tcp", serverSettings.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Printf("Failed to create zap logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	cluster := gocql.NewCluster("scylladb:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to create scylladb session: %v", err)
	}
	defer session.Close()
	log.Println("Scylladb session started")

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	log.Printf("Server listening on port %v", listen.Addr())
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	logger.Info("Inserting test")
}
