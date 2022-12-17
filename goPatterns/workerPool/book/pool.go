package book

import (
	"context"
	"log"
)

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- &Object{id: i}
	}

	return &p
}

type Object struct {
	id int
}

func (o *Object) Do(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		log.Printf("id %v do work", o.id)
	}
}
