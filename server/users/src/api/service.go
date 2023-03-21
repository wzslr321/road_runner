package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Service interface {
	handleGetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	handleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	handleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	handleRegisterUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	handleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	handleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error)
}
