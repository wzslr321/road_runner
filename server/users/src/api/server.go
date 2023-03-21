package api

import (
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Server struct {
	service Service
	pb.UnimplementedUsersServer
}

func NewServer(service Service) *Server {
	s := &Server{
		service: service,
	}
	return s
}
