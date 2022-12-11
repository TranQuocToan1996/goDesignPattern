package behavioralPatterns

// Although it may mirror all the methods declared in the context,
// aim only for those that may contain state-specific behavior

type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
