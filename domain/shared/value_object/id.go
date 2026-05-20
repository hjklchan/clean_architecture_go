package value_object

import "github.com/google/uuid"

type Id struct {
	value uuid.UUID
}

func NewId() Id {
	return Id{
		value: uuid.New(),
	}
}

func NewIdFromBytes(value [16]byte) Id {
	return Id{
		value: value,
	}
}

func (i Id) Value() uuid.UUID {
	return uuid.UUID(i.value)
}

func (i Id) BytesValue() [16]byte {
	return i.value
}

func (i Id) EqualTo(other Id) bool {
	return i.value == other.value
}
