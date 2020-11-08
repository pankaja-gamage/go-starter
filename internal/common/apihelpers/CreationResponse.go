package apihelpers

import (
	uuid "github.com/satori/go.uuid"
)

type CreationResponse struct {
	ID      uuid.UUID
	Message string
	Status  int
}
