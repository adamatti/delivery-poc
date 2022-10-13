package products

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

func TestListProductsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
  w := httptest.NewRecorder()
	
	listProductsHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestCreateProductHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/products", getTestdata(t, "testdata/create.json"))
  w := httptest.NewRecorder()
	
	createProductHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestUpdateProductHandler(t *testing.T) {
	product:= Product{Name: pointer("New Product")}
	product = *product.insert()
	assert.NotNil(t, product.Id)

	req := httptest.NewRequest(http.MethodPost, "/products", getTestdata(t, "testdata/update.json"))
	req = mux.SetURLVars(req, map[string]string{"id": product.Id.String()})
  w := httptest.NewRecorder()
	
	updateProductHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	if (!assert.Equal(t, w.Code, http.StatusOK)) {
		return
	}

	product = *findById(product.Id.String())
	assert.NotNil(t, product)
	assert.Equal(t, "updated", *product.Name)
}

func TestDeleteProductHandler(t *testing.T) {
	product:= Product{Name: pointer("New Product")}
	product = *product.insert()
	assert.NotNil(t, product.Id)

	req := httptest.NewRequest(http.MethodDelete, "/products", getTestdata(t, "testdata/update.json"))
	req = mux.SetURLVars(req, map[string]string{"id": product.Id.String()})
  w := httptest.NewRecorder()
	
	deleteProductHandler(w, req)
	
	res := w.Result()
  defer res.Body.Close()

	if (!assert.Equal(t, w.Code, http.StatusAccepted)) {
		return
	}

	productRef:= findById(product.Id.String())
	assert.Nil(t, productRef)
}

func getTestdata(t *testing.T, file string) io.ReadCloser {
	t.Helper()

	content, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	return io.NopCloser(bytes.NewReader(content))
}
