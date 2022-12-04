package behavioralPatterns

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}
