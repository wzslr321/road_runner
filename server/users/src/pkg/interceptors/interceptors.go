package interceptors

import (
	"context"
	"github.com/wzslr321/road_runner/server/users/src/pkg/auth"
	"github.com/wzslr321/road_runner/server/users/src/pkg/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

type InterceptorManager struct {
	metr   metrics.Metrics
	tokens auth.Jwt
}

func NewInterceptorManager(metr metrics.Metrics) *InterceptorManager {
	return &InterceptorManager{metr: metr}
}

func (im *InterceptorManager) Metrics(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	var status = http.StatusOK
	if err != nil {
		// TODO
	}
	im.metr.ObserveResponseTime(status, info.FullMethod, info.FullMethod, time.Since(start).Seconds())
	im.metr.IncHits(status, info.FullMethod, info.FullMethod)

	return resp, err
}

func (im *InterceptorManager) EnsureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, auth.ErrMissingMetadata
	}
	if !im.tokens.ValidateToken(md["authorization"]) {
		return nil, auth.ErrInvalidToken
	}
	return handler(ctx, req)
}
