package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/rides/src/proto-gen"
)

type Service interface {
	handleFindRide(req *pb.FindRideRequest, rides pb.Rides_FindRideServer) error
	handleStartRide(ctx context.Context, req *pb.StartRideRequest) (*pb.StartRideResponse, error)
	handleCancelRide(ctx context.Context, req *pb.CancelRideRequest) (*pb.CancelRideResponse, error)
	handleEndRide(ctx context.Context, req *pb.EndRideRequest) (*pb.EndRideResponse, error)
	handleCreateRide(ctx context.Context, req *pb.CreateRideRequest) (*pb.CreateRideResponse, error)
}
