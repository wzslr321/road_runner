package main

import (
	"crypto/tls"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wzslr321/road_runner/server/users/src/api"
	"github.com/wzslr321/road_runner/server/users/src/pkg/interceptors"
	"github.com/wzslr321/road_runner/server/users/src/pkg/metrics"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"net/http"
)

func main() {
	ms, err := metrics.Create("0.0.0.0:7070", "users")
	intercs := interceptors.NewInterceptorManager(ms)

	service := api.NewUsersService()
	service = api.NewLoggingService(service)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	cert, err := tls.LoadX509KeyPair("./cert/server_cert.pem", "./cert/server_key.pem")
	if err != nil {
		log.Fatalf("Failed to load cert: %v", err)
	}

	server := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.KeepaliveParams(keepalive.ServerParameters{}),
		grpc.UnaryInterceptor(intercs.Metrics),
		grpc.ChainUnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
		// grpc.ChainUnaryInterceptor(intercs.EnsureValidToken),
	)

	pb.RegisterUsersServer(server, api.NewServer(service))
	grpcprometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())
	server.Serve(listener)
}
