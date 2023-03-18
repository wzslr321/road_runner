package storage

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"log"
)

type Scylla struct {
	cluster *gocql.ClusterConfig
}

func New() *Scylla {
	cluster := gocql.NewCluster("scylladb:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	log.Println("Scylla cluster successfully created")

	return &Scylla{cluster: cluster}
}

func (s *Scylla) GetUser(username string) (*pb.GetUserResponse, error) {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	var users []*pb.GetUserResponse
	q := fmt.Sprintf("SELECT * FROM users.users WHERE username = '%s' ALLOW FILTERING", username)
	err = session.Query(q, nil).Select(&users)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return users[0], nil
}

func (s *Scylla) RegisterUser(user *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return &pb.RegisterUserResponse{Success: true}, nil
}

func (s *Scylla) LoginUser(user *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return &pb.LoginUserResponse{Success: true}, nil
}

func (s *Scylla) UpdateUser(user *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{Success: true}, nil
}

func (s *Scylla) LogoutUser(user *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error) {
	return &pb.LogoutUserResponse{Success: true}, nil
}
