package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/rides/src/proto-gen"
)

func (s *RidesService) handleFindRide(req *pb.FindRideRequest, rides pb.Rides_FindRideServer) error {
	panic("implement me")
}

func (s *RidesService) handleStartRide(ctx context.Context, req *pb.StartRideRequest) (*pb.StartRideResponse, error) {
	panic("implement me")
}

func (s *RidesService) handleCancelRide(ctx context.Context, req *pb.CancelRideRequest) (*pb.CancelRideResponse, error) {
	panic("implement me")
}

func (s *RidesService) handleEndRide(ctx context.Context, req *pb.EndRideRequest) (*pb.EndRideResponse, error) {
	panic("implement me")
}

func (s *RidesService) handleCreateRide(ctx context.Context, req *pb.CreateRideRequest) (*pb.CreateRideResponse, error) {
	panic("implement me")
}
