package ticket

type TicketStatus byte

var (
	StatusPending    TicketStatus = 1
	StatusProcessing TicketStatus = 2
	StatusTesting    TicketStatus = 3
	StatusDone       TicketStatus = 4
)

func DefaultTicketStatus() TicketStatus {
	return StatusPending
}
