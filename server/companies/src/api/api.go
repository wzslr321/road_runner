package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/companies/src/proto-gen"
)

func (s *Server) InviteDriver(ctx context.Context, req *pb.InviteDriverRequest) (*pb.InviteDriverResponse, error) {
	return s.service.handleInviteDriver(ctx, req)
}

func (s *Server) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	return s.service.handleGetInfo(ctx, req)
}

func (s *Server) GetDrivers(ctx context.Context, req *pb.GetDriversRequest) (*pb.GetDriversResponse, error) {
	return s.service.handleGetDrivers(ctx, req)
}

func (s *Server) UpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest) (*pb.UpdateInfoResponse, error) {
	return s.service.handleUpdateInfo(ctx, req)
}

func (s *Server) RegisterCompany(ctx context.Context, req *pb.RegisterCompanyRequest) (*pb.RegisterCompanyResponse, error) {
	return s.service.handleRegisterCompany(ctx, req)
}

func (s *Server) RemoveDriver(ctx context.Context, req *pb.RemoveDriverRequest) (*pb.RemoveDriverResponse, error) {
	return s.service.handleRemoveDriver(ctx, req)
}
