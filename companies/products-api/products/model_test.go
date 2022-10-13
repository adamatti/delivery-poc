package products

import ("testing")

func TestProductStringForNullPointers(t *testing.T) {
	product:= Product{}
	result:=product.String()

	if (result == "") {
		t.Error("Product.String() should return a string")
	}
}