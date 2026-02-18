package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {

	mode := os.Getenv("ENV");

	host 		:= os.Getenv("DATABASE_HOST");
	username	:= os.Getenv("DATABASE_USERNAME");
	password	:= os.Getenv("DATABASE_PASSWORD");
	dbName		:= os.Getenv("DATABASE_NAME");
	port		:= os.Getenv("DATABASE_PORT");

	var dsn string;

	if mode == "prod" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Bangkok prepare_threshold=0",
		host, username, password, dbName, port);
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, username, password, dbName, port);
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	});

	if err != nil {
		log.Fatal("⚠️ Failed to connect to database: ", err);
	}

	log.Println("✅ Contect to database successfully");

	return db;
}