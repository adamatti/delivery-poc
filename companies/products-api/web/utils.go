package web

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func JsonResponse(w http.ResponseWriter, r *http.Request, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		log.Error("Error sending response", err)
	}
}

func GetRequestData(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	return err
}

func SendError (w http.ResponseWriter, r *http.Request, statusCode int, msg string) {
	obj := make(map[string]string)
	obj["message"] = msg

	w.WriteHeader(statusCode)
	JsonResponse(w, r, obj)
}