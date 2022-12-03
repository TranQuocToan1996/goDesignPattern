package structuralPatterns

import "fmt"

type HP struct {
}

func (p *HP) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
