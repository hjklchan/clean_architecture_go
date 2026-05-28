package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Training struct {
	id              uuid.UUID
	userId          uuid.UUID
	userName        string
	time            time.Time
	notes           string
	proposedNewTime time.Time
	moveProposedBy  uuid.UUID
	canceled        bool
}

func NewTraining(id, userId uuid.UUID, userName string, trainingTime time.Time) (*Training, error) {
	if userName == "" {
		return nil, errors.New("user name can't be empty")
	}
	if trainingTime.IsZero() {
		return nil, errors.New("invalid training time")
	}

	return &Training{
		id:       id,
		userId:   userId,
		userName: userName,
		time:     trainingTime,
	}, nil
}

func (t *Training) CanBeCanceledForFree() bool {
	return time.Until(t.time) >= time.Hour*24
}

func (t *Training) IsCanceled() bool {
	return t.canceled
}

func (t *Training) Cancel() error {
	if t.IsCanceled() {
		return errors.New("training is already canceled")
	}

	t.canceled = true

	return nil
}
