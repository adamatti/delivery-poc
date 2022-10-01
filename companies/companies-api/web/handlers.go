package web

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// FIXME make a better return
// FIXME check infrastructure here
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := make(map[string]string)
	res["status"] = "ok"

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Error("Error sending response", err)
	}
}
