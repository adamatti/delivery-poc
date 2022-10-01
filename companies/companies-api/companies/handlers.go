package companies

import (
	"net/http"	
	"github.com/adamatti/delivery-poc/companies/web"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type ListCompaniesResponse struct {
	Data []Company `json:"data"`
}

// FIXME add additional properties like count, current page, next page
func listCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	res := ListCompaniesResponse{Data: list()}
	web.JsonResponse(w, r, res)
}

func getCompanyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if (!ok) {
		web.SendError(w, r, http.StatusInternalServerError, "Unable to parse id")
		return
	}

	company:= findById(id)
	if company == nil {
		web.SendError(w, r, http.StatusNotFound, "Company not found")
		return
	}

	web.JsonResponse(w, r, company)
}

func createCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var company Company
	defer r.Body.Close()

	if err := web.GetRequestData(r, &company); err != nil {
		log.Error(err)
		web.SendError(w, r, http.StatusBadRequest, "Invalid company object provided")
		return
	}
	
	web.JsonResponse(w, r, company.insert())
}

func updateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if (!ok) {
		web.SendError(w, r, http.StatusInternalServerError, "Unable to parse id")
		return
	}

	company:= findById(id)
	if company == nil {
		web.SendError(w, r, http.StatusNotFound, "Company not found")
		return
	}

	var companyProvided Company
	defer r.Body.Close()

	if err := web.GetRequestData(r, &companyProvided); err != nil {
		log.Error(err)
		web.SendError(w, r, http.StatusBadRequest, "Invalid company object provided")
		return
	}
	web.JsonResponse(w, r, company.update(companyProvided))
}

func deleteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if (!ok) {
		web.SendError(w, r, http.StatusInternalServerError, "Unable to parse id")
		return
	}

	company:= findById(id)
	if company == nil {
		web.SendError(w, r, http.StatusNotFound, "Company not found")
		return
	}
	
	company.delete()
	w.WriteHeader(http.StatusAccepted)
	web.JsonResponse(w, r, new(interface {}))
}

