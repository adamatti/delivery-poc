package products

import (
	"testing"
	"os"
	"github.com/adamatti/delivery-poc/products/database"
	
	"github.com/stretchr/testify/assert"
)

func startDatabase() {
	// FIXME start a mongo for tests
	uri:= "mongodb://mongo:mongo@localhost:27017/products-test"
	database.StartDatabase(uri)
}

func TestMain(m *testing.M) {
	startDatabase()
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {
	products:=list()
	assert.NotNil(t, products, "Products should be an array")
}

func TestFindByIdFinding(t *testing.T) {
	product:=Product{ Name: pointer("Fake product")}
	productRef:=product.insert()
	assert.NotNil(t, productRef, "Product should be created")

	productRef = findById(productRef.Id.String())
	assert.NotNil(t, productRef, "Should find an product")
}

func TestFindByIdNotFinding(t *testing.T) {
	productRef := findById("invalid")
	assert.Nil(t, productRef, "Should return nil when not found")
}

func TestDelete(t *testing.T) {
	product:=Product{ Name: pointer("Fake Product")}
	productRef:=product.insert()
	assert.NotNil(t, productRef,"Product should be created")

	product.delete()

	productRef = findById(product.Id.String())
	assert.Nil(t, productRef,"Product should be deleted")
}

func TestInsert(t *testing.T) {
	product:=Product{ Name: pointer("Fake Product")}
	productRef:=product.insert()
	assert.NotNil(t, productRef, "Product should be created")
}

func TestUpdate(t *testing.T) {
	product:=Product{ Name: pointer("Fake Product")}
	productRef:=product.insert()
	assert.NotNil(t, productRef, "Product should be created")

	productRef = productRef.update(Product{Name: pointer("New name")})
	assert.Equal(t, *productRef.Name, "New name", "Name should be updated")
}