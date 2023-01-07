package infra

import (
	"app/ent"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*ent.Client, error) {
	var DB_URL string
	if _, exist := os.LookupEnv("DATABASE_URL"); !exist {
		DB_URL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	} else {
		DB_URL = os.Getenv("DATABASE_URL")
	}

	client, err := ent.Open("postgres", DB_URL)
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
	}

	return client, err
}
