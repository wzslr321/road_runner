package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/companies/src/proto-gen"
)

type Service interface {
	handleInviteDriver(ctx context.Context, req *pb.InviteDriverRequest) (*pb.InviteDriverResponse, error)
	handleGetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error)
	handleGetDrivers(ctx context.Context, req *pb.GetDriversRequest) (*pb.GetDriversResponse, error)
	handleUpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest) (*pb.UpdateInfoResponse, error)
	handleRegisterCompany(ctx context.Context, req *pb.RegisterCompanyRequest) (*pb.RegisterCompanyResponse, error)
	handleRemoveDriver(ctx context.Context, req *pb.RemoveDriverRequest) (*pb.RemoveDriverResponse, error)
}
