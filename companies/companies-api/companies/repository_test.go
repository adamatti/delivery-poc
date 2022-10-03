package companies

import (
	"testing"
	"os"
	"github.com/adamatti/delivery-poc/companies/database"
	
	"gorm.io/driver/sqlite"
	"github.com/stretchr/testify/assert"
)

func startDatabase() {
	entities:= database.Entities{
		Company{},
	}
	database.StartDatabase(entities, sqlite.Open("gorm.db"))
}

func TestMain(m *testing.M) {
	startDatabase()
	code := m.Run()
	os.Exit(code)
}

func TestList(t *testing.T) {
	companies:=list()
	assert.NotNil(t, companies, "Companies should be an array")
}

func TestFindByIdFinding(t *testing.T) {
	company:=Company{ Name: pointer("Fake Company")}
	companyRef:=company.insert()
	assert.NotNil(t, companyRef, "Company should be created")

	companyRef = findById(companyRef.Id.String())
	assert.NotNil(t, companyRef, "Should find an company")
}

func TestFindByIdNotFinding(t *testing.T) {
	companyRef := findById("invalid")
	assert.Nil(t, companyRef, "Should return nil when not found")
}

func TestDelete(t *testing.T) {
	company:=Company{ Name: pointer("Fake Company")}
	companyRef:=company.insert()
	assert.NotNil(t, companyRef,"Company should be created")

	company.delete()

	companyRef = findById(company.Id.String())
	assert.Nil(t, companyRef,"Company should be deleted")
}

func TestInsert(t *testing.T) {
	company:=Company{ Name: pointer("Fake Company")}
	companyRef:=company.insert()
	assert.NotNil(t, companyRef, "Company should be created")
}

func TestUpdate(t *testing.T) {
	company:=Company{ Name: pointer("Fake Company")}
	companyRef:=company.insert()
	assert.NotNil(t, companyRef, "Company should be created")

	companyRef = companyRef.update(Company{Name: pointer("New name")})
	assert.Equal(t, *companyRef.Name, "New name", "Name should be updated")
}