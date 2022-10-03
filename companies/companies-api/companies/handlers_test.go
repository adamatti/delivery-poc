package companies

import (
	"bytes"
	"io"	
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestListCompaniesHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/companies", nil)
  w := httptest.NewRecorder()
	
	listCompaniesHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestCreateCompanyHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/companies", getTestdata(t, "testdata/create.json"))
  w := httptest.NewRecorder()
	
	createCompanyHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestUpdateCompanyHandler(t *testing.T) {
	company:= Company{Name: pointer("New Company")}
	company = *company.insert()
	assert.NotNil(t, company.Id)

	req := httptest.NewRequest(http.MethodPost, "/companies", getTestdata(t, "testdata/update.json"))
	req = mux.SetURLVars(req, map[string]string{"id": company.Id.String()})
  w := httptest.NewRecorder()
	
	updateCompanyHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	if (!assert.Equal(t, w.Code, http.StatusOK)) {
		return
	}

	company = *findById(company.Id.String())
	assert.NotNil(t, company)
	assert.Equal(t, "updated", *company.Name)
}

func TestDeleteCompanyHandler(t *testing.T) {
	company:= Company{Name: pointer("New Company")}
	company = *company.insert()
	assert.NotNil(t, company.Id)

	req := httptest.NewRequest(http.MethodDelete, "/companies", getTestdata(t, "testdata/update.json"))
	req = mux.SetURLVars(req, map[string]string{"id": company.Id.String()})
  w := httptest.NewRecorder()
	
	deleteCompanyHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	if (!assert.Equal(t, w.Code, http.StatusAccepted)) {
		return
	}

	companyRef:= findById(company.Id.String())
	assert.Nil(t, companyRef)
}

func getTestdata(t *testing.T, file string) io.ReadCloser {
	t.Helper()

	content, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	return io.NopCloser(bytes.NewReader(content))
}
