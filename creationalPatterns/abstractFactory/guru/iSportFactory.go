package creationalPatterns

import "fmt"

// Abstract factory interface
type iSportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type Brand string

const (
	adidas Brand = "adidas"
	nike   Brand = "nike"
)

func GetSportsFactory(brand Brand) (iSportFactory, error) {
	if brand == adidas {
		return &Adidas{}, nil
	}

	if brand == nike {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
