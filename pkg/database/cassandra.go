package database

import (
	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

func NewCassandra(host string, port string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(host, port)
	cluster.Keyspace = "chat"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}
