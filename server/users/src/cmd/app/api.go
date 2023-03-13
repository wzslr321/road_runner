package main

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type grpcServer struct {
	service Service
	pb.UnimplementedUsersServer
}

func newGrpcServer() *grpcServer {
	s := &grpcServer{
		service: NewUserService(),
	}

	return s
}

func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserData, error) {
	return s.service.GetUser(ctx)
}
