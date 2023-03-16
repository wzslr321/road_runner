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
	defer logger.Sync()
	return &LoggingService{child: child, logger: logger}
}

func (s *LoggingService) handleGetUser(ctx context.Context) (user *pb.UserData, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntook=%v", user, err, time.Since(start)))

	}(time.Now())
	return s.child.handleGetUser(ctx)
}
