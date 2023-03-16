package api

import (
	"github.com/gocql/gocql"
	"github.com/wzslr321/road_runner/server/users/src/storage"
)

type UsersService struct {
	storage *gocql.ClusterConfig
}

func NewUsersService() Service {
	return &UsersService{
		storage: storage.New(),
	}
}
