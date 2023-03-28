package api

import "github.com/wzslr321/road_runner/server/rides/src/storage"

type RidesService struct {
	db storage.Database
}

func NewRidesService() Service {
	return &RidesService{
		db: storage.New(),
	}
}
