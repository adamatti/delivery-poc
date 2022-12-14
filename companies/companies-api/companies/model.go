package companies

import (
	"time"
	"fmt"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Company struct {
	Id   *uuid.UUID `gorm:"primaryKey" json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	CreatedAt *time.Time `gorm:"<-:create;autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:nano" json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (company Company) String() string {
	var companyId string
	if (company.Id != nil) {
		companyId = company.Id.String()
	}

	var companyName string
	if (company.Name != nil) {
		companyName = *company.Name
	}

	return fmt.Sprintf("[id: %s, name: %s]", companyId, companyName)
}