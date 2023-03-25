package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

func (s *UsersService) handleGetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user, err := s.db.GetUser(req)
	if err != nil {
		return &pb.GetUserResponse{Code: "404", Message: err.Error()}, nil
	}

	return user, nil
}
func (s *UsersService) handleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}
func (s *UsersService) handleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

// i dont like this naming but im not sure if create will be better, gotta consider TODO
func (s *UsersService) handleCreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.db.CreateUser(req)
	if err != nil {
		return &pb.CreateUserResponse{Code: "500", Message: err.Error()}, nil
	}

	return user, nil
}

func (s *UsersService) handleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return nil, nil
}
func (s *UsersService) handleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error) {
	return nil, nil
}
