package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adamatti/delivery-poc/companies/config"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	readTimeout  = 10 // how long it takes to read the request
	writeTimeout = 60 // total time it will take to write the reques
)

type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunc     http.HandlerFunc
}

type Routes []Route

var (
	serverRoutes = Routes {
		Route{
			Name:        "Status",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: HealthCheckHandler,
		},
	}
	cfg = config.Instance
)

func getServerObj(router http.Handler) *http.Server {
	if cfg.ServicePort == 0 {
		log.Error("Service port not given")
		return nil
	}
	listenAddr := fmt.Sprintf(":%d", cfg.ServicePort)

	s := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	return s
}

func newRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func StartServer() {
	router := newRouter(serverRoutes)
	s := getServerObj(router)

	log.Infof("Starting up on %d...", cfg.ServicePort)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server Error: %v", err)
	}
	log.Info("Shutting down")
}