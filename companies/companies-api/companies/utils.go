package companies

import (
	"time"
	"github.com/google/uuid"
)

// FIXME review if it is needed
func pointer[O string | time.Time | uuid.UUID](o O) *O {
	return &o
}
