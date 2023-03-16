package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"log"
)

func (s *UsersService) handleGetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	session, err := s.storage.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	log.Printf("Scylla session sucessfully started")

	return &pb.GetUserResponse{
		Id:       "1",
		Email:    "yes@op.pl",
		Password: "1234",
		Username: "john",
	}, nil
}
func (s *UsersService) handleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}
func (s *UsersService) handleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}
func (s *UsersService) handleRegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return nil, nil
}
func (s *UsersService) handleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return nil, nil
}
func (s *UsersService) handleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error) {
	return nil, nil
}
