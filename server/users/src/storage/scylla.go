package storage

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"log"
)

type Scylla struct {
	cluster *gocql.ClusterConfig
}

func New() Scylla {
	cluster := gocql.NewCluster("scylladb:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	log.Println("Scylla cluster successfully created")

	return Scylla{cluster: cluster}
}

func (s *Scylla) GetUser(username string) (*pb.GetUserResponse, error) {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	err = session.ExecStmt(fmt.Sprintf(
		`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`,
		"users",
	))
	if err != nil {
		return nil, nil
	}

	err = session.ExecStmt(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users(
		id uuid PRIMARY KEY,
		username text,
		email text,
		)`, "users"))
	if err != nil {
		return nil, nil
	}

	insertUser := qb.Insert(fmt.Sprintf("%s.users", "users")).
		Columns("id", "username", "email").Query(session)
	insertUser.BindStruct(&pb.GetUserResponse{
		Id:       "35abf7b1-d847-4bbd-969e-487ddfb22d11",
		Username: "123",
		Email:    "123",
	})
	if err := insertUser.ExecRelease(); err != nil {
		return nil, err
	}
	var users []*pb.GetUserResponse
	err = session.Query("SELECT * FROM users.users", nil).Select(&users)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", users)

	return users[0], nil
}
