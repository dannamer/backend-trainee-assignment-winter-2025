package domain

import (
	"github.com/google/uuid"
)

type Merch struct {
	ID    uuid.UUID
	Item  string
	Price int
}
