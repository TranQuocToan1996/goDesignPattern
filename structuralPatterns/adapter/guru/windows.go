package structuralPatterns

import "fmt"

// Windows does not compatible with the computer interface
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
