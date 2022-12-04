package structuralPatterns

type TomatoTopping struct {
	pizza IPizza
}

// Concrete decorator
func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
