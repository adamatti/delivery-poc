package companies

import (
	"time"

	"github.com/adamatti/delivery-poc/companies/database"
	"github.com/google/uuid"
)

var companies = []*Company{
	buildByName("Pizzaria"),
	buildByName("Burguer King"),
	buildByName("Mc Donalds"),
}

func list() []Company {
	var companies []Company
	database.GetInstance().Find(&companies)
	return companies
}

func findById(id string) *Company {
	var company Company
	response:= database.GetInstance().Where("id = ?", id).First(&company)
	if response.RowsAffected == 0 {
		return nil
	}
	return &company
}

func (company Company) delete(){
	// This is a soft delete (gorm functionality)
	database.GetInstance().Delete(&company)
}

func buildByName(name string) *Company {
	return &Company{
		Id: pointer(uuid.New()),
		Name: &name,
		CreatedAt: pointer(time.Now()),
		UpdatedAt: pointer(time.Now()),
	}
}

func (company *Company) update(provided Company) *Company {	
	database.GetInstance().Model(&Company{}).
		Where("id = ?", company.Id.String()).
		Select("Name").
		Updates(provided)
	return findById(company.Id.String())
}

func (company *Company) insert() *Company {
	if company.Id == nil {
		company.Id = pointer(uuid.New())
	}
	database.GetInstance().Save(company)
	return company
}
