package main

import (
	"github.com/adamatti/delivery-poc/products/products"
	"github.com/adamatti/delivery-poc/products/config"
	"github.com/adamatti/delivery-poc/products/web"
	"github.com/adamatti/delivery-poc/products/database"

	log "github.com/sirupsen/logrus"
)

var cfg = config.Instance

func main() {
	log.Debug("Starting app")

	database.StartDatabase(cfg.MongoUrl)

	// Web layer
	appRoutes:= products.AppRoutes
	web.StartServer(appRoutes)
}
