package storage

import (
	"github.com/gocql/gocql"
	"log"
)

func New() *gocql.ClusterConfig {
	cluster := gocql.NewCluster("scylladb:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	log.Println("Scylla cluster successfully created")

	return cluster
}
