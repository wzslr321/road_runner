package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"log"
)

func (s *UsersService) handleGetUser(ctx context.Context) (*pb.UserData, error) {
	session, err := s.storage.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	log.Printf("Scylla session sucessfully started")

	return &pb.UserData{
		Id:       "1",
		Email:    "yes@op.pl",
		Password: "1234",
		Username: "john",
	}, nil
}
