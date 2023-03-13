package main

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Service interface {
	GetUser(ctx context.Context) (*pb.UserData, error)
}

type UserService struct{}

func NewUserService() Service {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context) (*pb.UserData, error) {
	return &pb.UserData{
		Id:       "1",
		Email:    "yes@op.pl",
		Password: "1234",
		Username: "john",
	}, nil
}
