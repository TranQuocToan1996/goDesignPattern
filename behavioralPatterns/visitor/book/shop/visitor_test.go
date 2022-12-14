package shop

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	products := make([]Visitable, 0)
	products = append(products, &Rice{
		Product: Product{
			Price: 32.0,
			Name:  "Some rice",
		}}, &Pasta{
		Product: Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}, &Fridge{
		Product: Product{
			Price: 50,
			Name:  "A fridge",
		}},
	)

	//Print the sum of prices
	priceVisitor := &PriceVisitor{}
	for _, p := range products {
		p.Accept(priceVisitor)
	}
	fmt.Printf("Total: %f\n", priceVisitor.Sum)
	//Print the products list
	nameVisitor := &NamePrinter{}
	for _, p := range products {
		p.Accept(nameVisitor)
	}
	fmt.Printf("\nProduct list:\n-------------\n%s",
		nameVisitor.ProductList)
}
