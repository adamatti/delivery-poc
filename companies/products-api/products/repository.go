package products

import (
	// "github.com/adamatti/delivery-poc/products/database"
	"github.com/google/uuid"
)

// FIXME implement
func list() []Product {
	var products []Product
	return products
}

// FIXME implement
func findById(id string) *Product {
	return nil
}

// FIXME implement
func (product Product) delete(){
	
}

// FIXME implement
func (product *Product) update(provided Product) *Product {	
	return findById(product.Id.String())
}

// FIXME implement
func (product *Product) insert() *Product {
	if product.Id == nil {
		product.Id = pointer(uuid.New())
	}
	return product
}
