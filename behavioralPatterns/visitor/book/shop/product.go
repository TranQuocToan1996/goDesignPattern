package shop

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}
func (p *Product) GetName() string {
	return p.Name
}

type Fridge struct {
	Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}
