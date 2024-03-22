package config

import (
	"database/sql"

	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

func DbPostgres() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=golang dbname=db_short_url sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func IntializeSetup() {
	db := DbPostgres()

	db.Exec("CREATE TABLE IF NOT EXISTS short_links (id VARCHAR(255) PRIMARY KEY, code VARCHAR(255), original_url VARCHAR(255), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
}

func DbRedis() *redis.Conn {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "admin",
		DB:       0,
	})

	return client.Conn()
}
