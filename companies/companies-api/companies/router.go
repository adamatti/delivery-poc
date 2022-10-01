package companies

import "github.com/adamatti/delivery-poc/companies/web"

var AppRoutes = web.Routes{
	web.Route{
		Name:        "List Companies",
		Method:      "GET",
		Pattern:     "/companies",
		HandlerFunc: listCompaniesHandler,
	},
	web.Route{
		Name:        "Get one Company by id",
		Method:      "GET",
		Pattern:     "/companies/{id:.*}",
		HandlerFunc: getCompanyHandler,
	},
	web.Route{
		Name:        "Create Company",
		Method:      "POST",
		Pattern:     "/companies",
		HandlerFunc: createCompanyHandler,
	},
	web.Route{
		Name:        "Update Company",
		Method:      "POST",
		Pattern:     "/companies/{id:.*}",
		HandlerFunc: updateCompanyHandler,
	},
	web.Route{
		Name:        "Delete Company",
		Method:      "DELETE",
		Pattern:     "/companies/{id:.*}",
		HandlerFunc: deleteCompanyHandler,
	},
}