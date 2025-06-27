package main

import (
	"log"
	"time"

	"github.com/Karan2980/go-grpc-graphql-microservice/account"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Simple retry logic without external library
	var r account.Repository
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(2 * time.Second)
		}
	}

	if err != nil {
		log.Fatal("Failed to connect to database after retries:", err)
	}

	defer r.Close()
	log.Println("Listening on port 8080...")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
