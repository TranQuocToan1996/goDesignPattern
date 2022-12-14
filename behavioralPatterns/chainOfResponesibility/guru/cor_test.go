package behavioralPatterns

import "testing"

func TestCOR(t *testing.T) {
	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
