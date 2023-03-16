package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Service interface {
	handleGetUser(ctx context.Context) (*pb.UserData, error)
}
