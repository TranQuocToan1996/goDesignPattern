package creationalPatterns

import (
	"testing"
)

func TestMethodFactory(t *testing.T) {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	generalTestGun(t, ak47)
	generalTestGun(t, musket)
}

func generalTestGun(t *testing.T, g iGun) {
	if len(g.getName()) == 0 {
		t.Fatal("must got name")
	}
	g.setPower(3)
	if g.getPower() != 3 {
		t.Error("not match power")
	}
}
