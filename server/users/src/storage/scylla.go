package storage

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/v2"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"github.com/wzslr321/road_runner/server/users/src/types"
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

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		// gotta handle it better
		log.Printf("Failed to create scylladb session: %v", err)
	}
	defer session.Close()

	err = tryToCreateKeyspace(&session)
	if err != nil {
		// gotta handle it better
		log.Printf("Failed to create keyspace: %v", err)
	}

	err = tryToCreateTable(&session)
	if err != nil {
		// gotta handle it better
		log.Printf("Failed to create table: %v", err)
	}

	return &Scylla{cluster: cluster}
}

// I think i should accept user model here and then convert it to proto
// but is it worth it tho?
func (s *Scylla) GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	var users []*types.User
	q := fmt.Sprintf("SELECT * FROM users.users WHERE username = '%s' ALLOW FILTERING", req.Username)
	err = session.Query(q, nil).Select(&users)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	user := users[0]
	return &pb.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Code:     "200",
		Message:  fmt.Sprintf("Successfully retrieved user with password %s", user.Password),
	}, nil
}

// make unique or sth
func (s *Scylla) CreateUser(req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	uid := uuid.Must(uuid.NewRandom()).String()

	q := fmt.Sprintf("INSERT INTO users.users (id, email, username, password) VALUES ('%s', '%s', '%s', '%s')", uid, req.Email, req.Username, req.Password)
	err = session.Query(q, nil).Exec()
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Code: "200", Message: "Successfully created a new user"}, nil
}

func (s *Scylla) UpdateUser(req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{Success: true}, nil
}

func (s *Scylla) DeleteUser(req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{Success: true}, nil
}

func tryToCreateKeyspace(session *gocqlx.Session) error {
	q := "CREATE KEYSPACE IF NOT EXISTS users WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}"
	err := session.Query(q, nil).Exec()
	if err != nil {
		return err
	}
	return nil
}

func tryToCreateTable(session *gocqlx.Session) error {
	q := "CREATE TABLE IF NOT EXISTS users.users (id text, email text, username text, password text, PRIMARY KEY (id))"
	err := session.Query(q, nil).Exec()
	if err != nil {
		return err
	}
	return nil
}
