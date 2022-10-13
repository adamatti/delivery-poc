package products

import (
	"time"
	"fmt"
	"github.com/google/uuid"
)

type Product struct {
	Id   *uuid.UUID `json:"id,omitempty"`
	CompanyId *string `json:"companyId,omitempty"`
	Name *string `json:"name,omitempty"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"-"`
}

func (product Product) String() string {
	var productId string
	if (product.Id != nil) {
		productId = product.Id.String()
	}

	var productName string
	if (product.Name != nil) {
		productName = *product.Name
	}

	return fmt.Sprintf("[id: %s, name: %s]", productId, productName)
}