package companies

import ("testing")

func TestCompanyStringForNullPointers(t *testing.T) {
	company:= Company{}
	result:=company.String()

	if (result == "") {
		t.Error("Company.String() should return a string")
	}
}