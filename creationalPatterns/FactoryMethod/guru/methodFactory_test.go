package creationalPatterns

import (
	"testing"
)

func TestMethodFactory(t *testing.T) {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(t, ak47)
	printDetails(t, musket)
}

func printDetails(t *testing.T, g iGun) {
	if len(g.getName()) == 0 {
		t.Fatal("must got name")
	}
	g.setPower(3)
	if g.getPower() != 3 {
		t.Error("not match power")
	}
}
