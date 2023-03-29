package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/companies/src/proto-gen"
)

func (s *CompaniesService) handleInviteDriver(ctx context.Context, req *pb.InviteDriverRequest) (*pb.InviteDriverResponse, error) {
	panic("implement me")
}

func (s *CompaniesService) handleGetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	panic("implement me")
}

func (s *CompaniesService) handleGetDrivers(ctx context.Context, req *pb.GetDriversRequest) (*pb.GetDriversResponse, error) {
	panic("implement me")
}

func (s *CompaniesService) handleUpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest) (*pb.UpdateInfoResponse, error) {
	panic("implement me")
}

func (s *CompaniesService) handleRegisterCompany(ctx context.Context, req *pb.RegisterCompanyRequest) (*pb.RegisterCompanyResponse, error) {
	panic("implement me")
}

func (s *CompaniesService) handleRemoveDriver(ctx context.Context, req *pb.RemoveDriverRequest) (*pb.RemoveDriverResponse, error) {
	panic("implement me")
}
