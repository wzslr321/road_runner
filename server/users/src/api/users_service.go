package api

import (
	"github.com/wzslr321/road_runner/server/users/src/storage"
)

type UsersService struct {
	db storage.Database
}

func NewUsersService() Service {
	return &UsersService{
		db: storage.New(),
	}
}
