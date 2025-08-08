package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"study-gin/internal/config"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := config.NewConfig().DNS()

	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}

	DB.SetMaxIdleConns(3)                   //so ket noi nhan roi toi da
	DB.SetMaxOpenConns(30)                  //so ket noi toi da
	DB.SetConnMaxLifetime(30 * time.Minute) //dong ket loi sau 30 phut
	DB.SetConnMaxIdleTime(5 * time.Minute)  //dong ket loi nhan roi sau 5 phut

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		DB.Close()
		return fmt.Errorf("DB ping error: %w", err)
	}

	log.Println("Connected")

	return nil
}
