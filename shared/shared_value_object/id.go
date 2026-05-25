package shared_value_object

import "github.com/google/uuid"

type Id uuid.UUID

func NewId() Id {
	return Id(uuid.New())
}

func (i Id) GetId() uuid.UUID {
	return uuid.UUID(i)
}

func (i Id) GetStringId() string {
	return i.GetId().String()
}

func (i Id) EqualTo(target Id) bool {
	return i.GetId() == target.GetId()
}
