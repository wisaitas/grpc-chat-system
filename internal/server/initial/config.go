package initial

import (
	"fmt"
	"log"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wisaitas/grpc-chat-system/internal/server"
	"github.com/wisaitas/grpc-chat-system/pkg/database"
)

type config struct {
	Postgres  *pgxpool.Pool
	Cassandra *gocql.Session
}

func newConfig() *config {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		server.Config.Postgres.Host,
		server.Config.Postgres.User,
		server.Config.Postgres.Password,
		server.Config.Postgres.DBName,
		server.Config.Postgres.Port,
	)

	dbClient, err := database.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}

	cassandraClient, err := database.NewCassandra(server.Config.Cassandra.Host, server.Config.Cassandra.Port)
	if err != nil {
		log.Fatalf("failed to create cassandra connection: %v", err)
	}

	return &config{
		Postgres:  dbClient,
		Cassandra: cassandraClient,
	}
}
