package initial

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wisaitas/grpc-chat-system/internal/server"
	"github.com/wisaitas/grpc-chat-system/pkg/database"
)

type config struct {
	Postgres *pgxpool.Pool
}

func newConfig() *config {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		server.Config.Postgres.Host,
		server.Config.Postgres.User,
		server.Config.Postgres.Password,
		server.Config.Postgres.DBName,
		server.Config.Postgres.Port,
	)

	return &config{
		Postgres: database.NewPostgres(dsn),
	}
}
