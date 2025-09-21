package initial

import (
	"fmt"

	"github.com/wisaitas/grpc-chat-system/internal/server"
	"github.com/wisaitas/grpc-chat-system/pkg/database"
	"gorm.io/gorm"
)

type config struct {
	postgres *gorm.DB
}

func newConfig() *config {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		server.Config.Postgres.Host,
		server.Config.Postgres.User,
		server.Config.Postgres.Password,
		server.Config.Postgres.DBName,
		server.Config.Postgres.Port,
	)

	return &config{
		postgres: database.NewPostgres(dsn),
	}
}
