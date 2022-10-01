package companies

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	Id   *uuid.UUID `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}