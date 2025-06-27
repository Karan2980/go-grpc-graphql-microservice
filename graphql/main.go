package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL string `envconfig:"Account_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	s, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL,cfg.OrderURL)
	if err != nil{
		log.Fatal(err)
	}

	http.Handle("/grpahql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("karan","/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}