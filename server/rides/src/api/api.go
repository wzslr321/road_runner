package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/rides/src/proto-gen"
)

func (s *Server) FindRide(req *pb.FindRideRequest, rides pb.Rides_FindRideServer) error {
	return s.service.handleFindRide(req, rides)
}

func (s *Server) StartRide(ctx context.Context, req *pb.StartRideRequest) (*pb.StartRideResponse, error) {
	return s.service.handleStartRide(ctx, req)
}

func (s *Server) CancelRide(ctx context.Context, req *pb.CancelRideRequest) (*pb.CancelRideResponse, error) {
	return s.service.handleCancelRide(ctx, req)
}

func (s *Server) EndRide(ctx context.Context, req *pb.EndRideRequest) (*pb.EndRideResponse, error) {
	return s.service.handleEndRide(ctx, req)
}

func (s *Server) CreateRide(ctx context.Context, req *pb.CreateRideRequest) (*pb.CreateRideResponse, error) {
	return s.service.handleCreateRide(ctx, req)
}
