package main

import (
	"github.com/adamatti/delivery-poc/companies/companies"
	"github.com/adamatti/delivery-poc/companies/config"
	"github.com/adamatti/delivery-poc/companies/web"

	log "github.com/sirupsen/logrus"
)

var cfg = config.Instance

func main() {
	log.Debug("Starting app")

	appRoutes:= companies.AppRoutes

	web.StartServer(appRoutes)
}
