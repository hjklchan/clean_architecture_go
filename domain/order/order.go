package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id uuid.UUID

	lines []any

	createdAt time.Time
	updatedAt time.Time
}
