package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgres(
	dsn string,
) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info), // แสดง SQL logs (development)
		PrepareStmt:            true,                                // เปิด prepared statements เพื่อเพิ่มประสิทธิภาพ
		SkipDefaultTransaction: true,                                // ปิด default transaction เพื่อเพิ่มประสิทธิภาพ
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)           // จำนวน idle connections
	sqlDB.SetMaxOpenConns(100)          // จำนวน active connections
	sqlDB.SetConnMaxLifetime(time.Hour) // อายุของ connection

	return db
}
