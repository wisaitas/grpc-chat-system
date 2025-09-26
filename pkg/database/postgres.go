package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %v", err)
	}

	config.MaxConns = 100
	config.MinConns = 10
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return pool, nil
}

// package database

// import (
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func NewPostgres(
// 	dsn string,
// ) *gorm.DB {
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger:                 logger.Default.LogMode(logger.Info), // แสดง SQL logs (development)
// 		PrepareStmt:            true,                                // เปิด prepared statements เพื่อเพิ่มประสิทธิภาพ
// 		SkipDefaultTransaction: true,                                // ปิด default transaction เพื่อเพิ่มประสิทธิภาพ
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	sqlDB.SetMaxIdleConns(10)           // จำนวน idle connections
// 	sqlDB.SetMaxOpenConns(100)          // จำนวน active connections
// 	sqlDB.SetConnMaxLifetime(time.Hour) // อายุของ connection

// 	return db
// }
