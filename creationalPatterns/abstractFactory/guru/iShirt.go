package creationalPatterns

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) GetLogo() string {
	return s.logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) GetSize() int {
	return s.size
}
