package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DbConnection *sql.DB
	once         sync.Once
)

func InitPostgresDB() *sql.DB {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", Env("DB_HOST"), Env("DB_USER"), Env("DB_PASS"), Env("DB_NAME"), Env("DB_PORT"))
		DbConnection, err = sql.Open(Env("DB_DRIVER"), dsn)
		if err != nil {
			log.Fatalf("unable to connect with database: %v", err)
		}
	})
	return DbConnection
}
