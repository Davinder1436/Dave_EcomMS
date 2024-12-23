package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl string `envconfig:"ACCOUNT_SERVICE_URL"`
	ProductUrl string `envconfig:"PRODUCT_SERVICE_URL"`
	OrderUrl   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	s, err := NewGraphServer(cfg.AccountUrl, cfg.ProductUrl, cfg.OrderUrl)
	if err != nil {
		panic(err)
	}
	http.Handle("/graphql", handler.New(s.ToExecuteableSchema()))
	http.Handle("/playground", playground.Handler("dave", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
