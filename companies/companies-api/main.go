package main

import (
	"github.com/adamatti/delivery-poc/companies/companies"
	"github.com/adamatti/delivery-poc/companies/config"
	"github.com/adamatti/delivery-poc/companies/web"
	"github.com/adamatti/delivery-poc/companies/database"

	"gorm.io/driver/postgres"
	log "github.com/sirupsen/logrus"
)

var cfg = config.Instance

func main() {
	log.Debug("Starting app")

	// Database
	config:= postgres.Config{
		// FIXME move string to config
		// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
		DSN: "host=localhost user=companies password=companies dbname=companies port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	entities:= database.Entities{
		companies.Company{},
	}
	database.StartDatabase(entities, postgres.New(config))

	// Web layer
	appRoutes:= companies.AppRoutes
	web.StartServer(appRoutes)
}
