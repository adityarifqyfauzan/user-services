package errors

type BadRequestError struct {
	s string
}

func NewBadRequestError(s string) error {
	return &BadRequestError{s: s}
}

func (e *BadRequestError) Error() string {
	return e.s
}
