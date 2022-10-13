package products

import (
	"net/http"	
	"github.com/adamatti/delivery-poc/products/web"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type ListProductsResponse struct {
	Data []Product `json:"data"`
}

// FIXME add additional properties like count, current page, next page
func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	res := ListProductsResponse{Data: list()}
	web.JsonResponse(w, r, res)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if (!ok) {
		web.SendError(w, r, http.StatusInternalServerError, "Unable to parse id")
		return
	}

	product:= findById(id)
	if product == nil {
		web.SendError(w, r, http.StatusNotFound, "Product not found")
		return
	}

	web.JsonResponse(w, r, product)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	defer r.Body.Close()

	if err := web.GetRequestData(r, &product); err != nil {
		log.Error(err)
		web.SendError(w, r, http.StatusBadRequest, "Invalid product object provided")
		return
	}
	
	web.JsonResponse(w, r, product.insert())
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if (!ok) {
		web.SendError(w, r, http.StatusInternalServerError, "Unable to parse id")
		return
	}

	product:= findById(id)
	if product == nil {
		web.SendError(w, r, http.StatusNotFound, "Product not found")
		return
	}

	var productProvided Product
	defer r.Body.Close()

	if err := web.GetRequestData(r, &productProvided); err != nil {
		log.Error(err)
		web.SendError(w, r, http.StatusBadRequest, "Invalid product object provided")
		return
	}
	web.JsonResponse(w, r, product.update(productProvided))
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if (!ok) {
		web.SendError(w, r, http.StatusInternalServerError, "Unable to parse id")
		return
	}

	product:= findById(id)
	if product == nil {
		web.SendError(w, r, http.StatusNotFound, "Product not found")
		return
	}
	
	product.delete()
	w.WriteHeader(http.StatusAccepted)
	web.JsonResponse(w, r, new(interface {}))
}

