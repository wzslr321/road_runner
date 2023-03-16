package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Service interface {
	handleGetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error)
	handleUpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	handleDeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	handleRegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	handleLoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	handleLogoutUser(ctx context.Context, request *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error)
}
