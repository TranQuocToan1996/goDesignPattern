package structuralPatterns


// Abstraction
type Computer interface {
	Print()
	SetPrinter(Printer)
}
