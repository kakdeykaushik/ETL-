package db

import (
	"etl/utils"
	"log"

	"github.com/gocql/gocql"
)

// var Session *gocql.Session

// todo - handle exploitation of New Session
func GetSession() *gocql.Session {
	// var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "etl"

	session, err := cluster.CreateSession()
	if err != nil {
		utils.FatalErr(err, "error connecting to cassandra")
	}
	log.Println("cassandra well initialized")
	return session
}
