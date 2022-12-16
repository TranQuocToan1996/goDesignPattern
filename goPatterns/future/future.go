package future

type (
	SuccessFunc func(string)

	FailFunc func(error)

	ExecuteStringFunc func() (string, error)

	MaybeString struct{}
)

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	return nil
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	return nil
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
}
