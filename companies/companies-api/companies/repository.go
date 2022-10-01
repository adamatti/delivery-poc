package companies

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var companies = []*Company{
	buildByName("Pizzaria"),
	buildByName("Burguer King"),
	buildByName("Mc Donalds"),
}

func findById(id string) *Company {
	for _ , c  := range companies {
		if (c.Id.String() == id) {
			return c
		}
	}
	return nil
}

func (company Company) delete(){
	newList := []*Company{}
	for _ , c  := range companies {
		if (c.Id.String() != company.Id.String()) {
			newList = append(newList, c)
		}
	}
	companies = newList
}

func buildByName(name string) *Company {
	return &Company{
		Id: pointer(uuid.New()),
		Name: &name,
		CreatedAt: pointer(time.Now()),
		UpdatedAt: pointer(time.Now()),
	}
}

func (company *Company) update(provided Company) Company {
	company.Name = provided.Name
	company.UpdatedAt = pointer(time.Now())

	return *company
}

func (company Company) insert() Company {
	if company.Id == nil {
		company.Id = pointer(uuid.New())
	}
	company.CreatedAt = pointer(time.Now())
	company.UpdatedAt = pointer(time.Now())

	companies = append(companies, &company)
	return company
}

func (company Company) String() string {
	return fmt.Sprintf("%s - %s", company.Id.String(), *company.Name)
}