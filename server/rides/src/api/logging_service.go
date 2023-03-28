package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/rides/src/proto-gen"
	"go.uber.org/zap"
	"time"
)

type LoggingService struct {
	logger *zap.Logger
	child  Service
}

func NewLoggingService(child Service) Service {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck
	return &LoggingService{child: child, logger: logger}
}

func (s *LoggingService) handleFindRide(req *pb.FindRideRequest, rides pb.Rides_FindRideServer) (err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleFindRide(req, rides)
}

func (s *LoggingService) handleStartRide(ctx context.Context, req *pb.StartRideRequest) (ride *pb.StartRideResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")

	}(time.Now())
	return s.child.handleStartRide(ctx, req)
}

func (s *LoggingService) handleCancelRide(ctx context.Context, req *pb.CancelRideRequest) (ride *pb.CancelRideResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleCancelRide(ctx, req)
}

func (s *LoggingService) handleEndRide(ctx context.Context, req *pb.EndRideRequest) (ride *pb.EndRideResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleEndRide(ctx, req)
}

func (s *LoggingService) handleCreateRide(ctx context.Context, req *pb.CreateRideRequest) (ride *pb.CreateRideResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleCreateRide(ctx, req)
}
