package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
	Port       string `envconfig:"PORT" default:"8080"`
}

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Set default URLs if not provided
	if cfg.AccountURL == "" {
		cfg.AccountURL = "localhost:8080"
	}
	if cfg.CatalogURL == "" {
		cfg.CatalogURL = "localhost:8081"
	}
	if cfg.OrderURL == "" {
		cfg.OrderURL = "localhost:8082"
	}

	log.Printf("Connecting to services:")
	log.Printf("Account: %s", cfg.AccountURL)
	log.Printf("Catalog: %s", cfg.CatalogURL)
	log.Printf("Order: %s", cfg.OrderURL)

	server, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)
	if err != nil {
		log.Fatal("Failed to create GraphQL server:", err)
	}

	srv := handler.NewDefaultServer(server.ToExecutableSchema())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL server running on http://localhost:%s/", cfg.Port)
	log.Printf("GraphQL playground available at http://localhost:%s/", cfg.Port)
	
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
