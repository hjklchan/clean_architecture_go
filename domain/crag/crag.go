package crag

import (
	"time"

	"github.com/google/uuid"
)

type Crag struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}

func NewCrag(name, desc, country string) Crag {
	return Crag{
		ID:        uuid.New(),
		Name:      name,
		Desc:      desc,
		Country:   country,
		CreatedAt: time.Now(),
	}
}
