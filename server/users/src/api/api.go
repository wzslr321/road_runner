package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserData, error) {
	return s.service.handleGetUser(ctx)
}
