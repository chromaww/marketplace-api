package base

import (
	"time"

	"github.com/google/uuid"
)

type ModelHasAudit struct {
	CreatedBy uuid.UUID  `json:"-"` // Hide metadata
	CreatedAt time.Time  `json:"created_at"`
	UpdatedBy *uuid.UUID `json:"-"` // Hide metadata
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedBy *uuid.UUID `json:"-"` // Hide metadata
	DeletedAt *time.Time `json:"deleted_at"`
}
