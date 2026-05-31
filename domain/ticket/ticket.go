package ticket

import (
	"errors"

	"github.com/google/uuid"
)

type Ticket struct {
	id uuid.UUID

	title       string
	description string
	assigneeId  uuid.UUID
	status      TicketStatus
}

var (
	ErrEmptyAssigneeId = errors.New("assignee id can not be empty")
	ErrEmptyTitle      = errors.New("title can not be empty")
)

func NewTicket(assigneeId uuid.UUID, title string, description string) (*Ticket, error) {
	if assigneeId.String() == "" {
		return nil, ErrEmptyAssigneeId
	}
	if title == "" {
		return nil, ErrEmptyTitle
	}

	return &Ticket{
		id:          uuid.New(),
		title:       title,
		description: description,
		assigneeId:  assigneeId,
		status:      StatusPending,
	}, nil
}
