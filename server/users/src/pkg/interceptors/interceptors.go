package interceptors

import (
	"context"
	"github.com/wzslr321/road_runner/server/users/src/pkg/metrics"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type InterceptorManager struct {
	metr metrics.Metrics
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
