package main

type AppConfig struct {
	AccountUrl string `envconfig:"ACCOUNT_SERVICE_URL`
	ProductUrl string `envconfig:"PRODUCT_SERVICE_URL`
	OrderUrl   string `envconfig:"ORDER_SERVICE_URL`
}
