package initial

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wisaitas/grpc-chat-system/internal/server"
	"github.com/wisaitas/grpc-chat-system/pkg/database"
)

type config struct {
	DB *pgxpool.Pool
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

	return &config{
		DB: dbClient,
	}
}
