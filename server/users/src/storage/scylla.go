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

func (s *Scylla) GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	var users []*pb.GetUserResponse
	q := fmt.Sprintf("SELECT * FROM users.users WHERE username = '%s' ALLOW FILTERING", req.Username)
	err = session.Query(q, nil).Select(&users)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return users[0], nil
}

func (s *Scylla) CreateUser(req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{Success: true}, nil
}

func (s *Scylla) UpdateUser(req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{Success: true}, nil
}

func (s *Scylla) DeleteUser(req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{Success: true}, nil
}
