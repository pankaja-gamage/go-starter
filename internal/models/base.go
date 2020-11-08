package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Model is common model structure
type Model struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
