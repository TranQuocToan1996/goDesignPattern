package creationalPatterns

import "fmt"

// Abstract factory interface
type ISportFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

type Brand string

const (
	adidas Brand = "adidas"
	nike   Brand = "nike"
)

func GetSportsFactory(brand Brand) (ISportFactory, error) {
	if brand == adidas {
		return &Adidas{}, nil
	}

	if brand == nike {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
