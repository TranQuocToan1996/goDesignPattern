package anhthi

type AsymEncryptionStrategy interface {
	Encrypt(data interface{}) (cipher []byte, err error)
}

type EllipticCurvestrategy struct{}
type RSA struct{}

func (strat *EllipticCurvestrategy) Encrypt(data interface{}) (cipher []byte, err error) {
	// some complex math
	return cipher, err
}

func (strat *RSA) Encrypt(data interface{}) (cipher []byte, err error) {
	// some complex math

	return cipher, err
}

func encryptMessage(msg string, strat AsymEncryptionStrategy) (cipher []byte, err error) {
	return strat.Encrypt(msg)
}

// usage
// msg := "this is a confidential message"
// cipher, err := encryptMessage(msg, encryptMessage(new(RSA)))
// cipher, err = encryptMessage(msg, encryptMessage(new(EllipticCurvestrategy)))
