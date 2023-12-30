package main

import (
	"log"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "etl"
	Session, err = cluster.CreateSession()
	if err != nil {
		fatalErr(err, "error connecting to cassandra")
	}
	log.Println("cassandra well initialized")
}
