package errors

type UnprocessableEntity struct {
	s string
}

func NewUnprocessabelEntity(s string) error {
	return &UnprocessableEntity{
		s: s,
	}
}

func (e *UnprocessableEntity) Error() string {
	return e.s
}
