package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T){
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
  w := httptest.NewRecorder()
	
	HealthCheckHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	assert.Equal(t, w.Code, http.StatusOK)
}