package creationalPatterns

type AK47 struct {
	Gun
}

func newAK47() iGun {
	return &AK47{
		Gun{
			name:  "AK47",
			power: 4,
		},
	}
}
