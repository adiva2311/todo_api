package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func NewDb() *sql.DB {
	dsn := "root:12345678@tcp(localhost:3306)/todo_api?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Database connected successfully!")

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}