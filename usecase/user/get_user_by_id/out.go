package get_user_by_id

type Output struct {
	ID    string
	Name  string
	Email string
}

type ErrorHandler interface {
	HandleError(error) error
}

type OutputPort interface {
	Present(Output) error

	ErrorHandler
}
