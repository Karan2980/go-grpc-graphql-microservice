package main

import (
	"log"
	"time"

	"github.com/Karan2980/go-grpc-graphql-microservice/catalog"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
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

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("Failed to connect to database:", err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, 8080))
}
