package products

import "github.com/adamatti/delivery-poc/products/web"

var AppRoutes = web.Routes{
	web.Route{
		Name:        "List Products",
		Method:      "GET",
		Pattern:     "/products",
		HandlerFunc: listProductsHandler,
	},
	web.Route{
		Name:        "Get one Product by id",
		Method:      "GET",
		Pattern:     "/products/{id:.*}",
		HandlerFunc: getProductHandler,
	},
	web.Route{
		Name:        "Create Product",
		Method:      "POST",
		Pattern:     "/products",
		HandlerFunc: createProductHandler,
	},
	web.Route{
		Name:        "Update Product",
		Method:      "POST",
		Pattern:     "/products/{id:.*}",
		HandlerFunc: updateProductHandler,
	},
	web.Route{
		Name:        "Delete Product",
		Method:      "DELETE",
		Pattern:     "/products/{id:.*}",
		HandlerFunc: deleteProductHandler,
	},
}