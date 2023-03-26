package api

import (
	"context"
	"fmt"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
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

// TODO make these message better, placeholders for now

func (s *LoggingService) handleGetUser(ctx context.Context, req *pb.GetUserRequest) (user *pb.GetUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntook=%v", user, err, time.Since(start)))

	}(time.Now())
	return s.child.handleGetUser(ctx, req)
}

func (s *LoggingService) handleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (user *pb.UpdateUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))

	}(time.Now())
	return s.child.handleUpdateUser(ctx, req)
}

func (s *LoggingService) handleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (user *pb.DeleteUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.handleDeleteUser(ctx, req)
}

func (s *LoggingService) handleCreateUser(ctx context.Context, req *pb.CreateUserRequest) (user *pb.CreateUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.handleCreateUser(ctx, req)
}

func (s *LoggingService) handleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (user *pb.LoginUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.handleLoginUser(ctx, req)
}

func (s *LoggingService) handleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (user *pb.LogoutUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.handleLogoutUser(ctx, req)
}
