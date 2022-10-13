package web

import (
	"net/http"
)

// FIXME make a better return
// FIXME check infrastructure here
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	obj := make(map[string]string)
	obj["status"] = "ok"

	JsonResponse(w, r, obj)
}
