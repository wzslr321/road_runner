package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/companies/src/proto-gen"
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

func (s *LoggingService) handleInviteDriver(ctx context.Context, req *pb.InviteDriverRequest) (resp *pb.InviteDriverResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleInviteDriver(ctx, req)
}

func (s *LoggingService) handleGetInfo(ctx context.Context, req *pb.GetInfoRequest) (resp *pb.GetInfoResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleGetInfo(ctx, req)
}

func (s *LoggingService) handleGetDrivers(ctx context.Context, req *pb.GetDriversRequest) (resp *pb.GetDriversResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleGetDrivers(ctx, req)
}

func (s *LoggingService) handleUpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest) (resp *pb.UpdateInfoResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleUpdateInfo(ctx, req)
}

func (s *LoggingService) handleRegisterCompany(ctx context.Context, req *pb.RegisterCompanyRequest) (resp *pb.RegisterCompanyResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleRegisterCompany(ctx, req)
}

func (s *LoggingService) handleRemoveDriver(ctx context.Context, req *pb.RemoveDriverRequest) (resp *pb.RemoveDriverResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info("IMPLEMENT LOGGER MAN")
	}(time.Now())
	return s.child.handleRemoveDriver(ctx, req)
}
