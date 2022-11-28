package creationalPatterns

type M4A1 struct {
	Gun
}

func newM4A1() iGun {
	return &M4A1{
		Gun{name: "M4A1", power: 3},
	}
}
