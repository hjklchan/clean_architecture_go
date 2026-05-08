package crag

import (
	"context"

	"github.com/google/uuid"
)

type CragRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (error, *Crag)
	GetAll(ctx context.Context) (error, []Crag)
	Add(ctx context.Context, crag Crag) error
	Update(ctx context.Context, crag Crag) error
	Delete(ctx context.Context, id uuid.UUID) error
}
