package creationalpatterns

import "testing"

func TestGetIntacce(t *testing.T) {
	counter1 := GetIntances()

	if counter1 == nil {
		t.Error("expect got counter")
	}

	expectedCounter := counter1

	currentCounter := counter1.AddOne()
	if currentCounter != 1 {
		t.Errorf("Expect got 1 but actually got %d \n", currentCounter)
	}

	counter2 := GetIntances()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
}
