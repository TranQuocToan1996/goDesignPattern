package structuralPatterns

import (
	"io"
	"testing"
)

func TestAdapterPrinter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}

	if adapter.PrintStored() != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", adapter.PrintStored())
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != msg {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
    io.Pipe()
}
