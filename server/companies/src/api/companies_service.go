package api

import "github.com/wzslr321/road_runner/server/companies/src/storage"

type CompaniesService struct {
	db storage.Database
}

func NewCompaniesService() Service {
	return &CompaniesService{
		db: storage.New(),
	}
}
