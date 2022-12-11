package shop

import "fmt"

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}
type Visitor interface {
	Visit(ProductInfoRetriever)
}
type Visitable interface {
	Accept(Visitor)
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}
