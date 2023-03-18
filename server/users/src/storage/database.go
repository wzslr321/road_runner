package storage

import pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"

type Database interface {
	GetUser(username string) (*pb.GetUserResponse, error)
	RegisterUser(user *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	LoginUser(user *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	UpdateUser(user *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	LogoutUser(user *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error)
}
