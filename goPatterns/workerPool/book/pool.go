package book

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

type Object struct {
}

func (o *Object) Do() {
	//TODO
}
