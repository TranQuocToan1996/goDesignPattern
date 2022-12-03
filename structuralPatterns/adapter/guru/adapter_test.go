package structuralPatterns

import "testing"

func TestAdapter(t *testing.T) {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windows := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windows,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
