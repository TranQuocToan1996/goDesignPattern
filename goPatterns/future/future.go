package future

type (
	SuccessFunc func(string)

	FailFunc func(error)

	ExecuteStringFunc func() (string, error)

	MaybeString struct {
		successFunc SuccessFunc
		failFunc    FailFunc
	}
)

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}
