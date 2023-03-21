package storage

import pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"

type Database interface {
	CreateUser(req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	UpdateUser(req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}
