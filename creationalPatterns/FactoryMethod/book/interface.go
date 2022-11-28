package creationalPatterns

import (
	"errors"
	"fmt"
)

const (
	Cash      = 1
	DebitCard = 2
)

type Paymethod interface {
	Pay(amount int64) string
}

func GetPaymentMethod(typ int) (Paymethod, error) {
	switch typ {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	}

	return nil, errors.New("not yet implementation")
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount int64) string {
	return fmt.Sprintf("%v paid using cash\n", amount)
}
func (c *DebitCardPM) Pay(amount int64) string {
	return fmt.Sprintf("%v paid using debit card\n", amount)
}
