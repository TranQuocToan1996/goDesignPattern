package creationalPatterns

import (
	"strings"
	"testing"
)

func TestPaymentMethod(t *testing.T) {
	methodCash, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("must has payment method")
	}

	result := methodCash.Pay(10000)
	if !strings.Contains(result, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}

}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Error("A payment method of type 'DebitCard' must exist")
	}
	msg := payment.Pay(15000)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	payment, err := GetPaymentMethod(20)
	if err == nil || payment != nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
