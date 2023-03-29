package api

import pb "github.com/wzslr321/road_runner/server/companies/src/proto-gen"

type Server struct {
	service Service
	pb.UnimplementedCompaniesServer
}

func NewServer(service Service) *Server {
	s := &Server{
		service: service,
	}
	return s
}
