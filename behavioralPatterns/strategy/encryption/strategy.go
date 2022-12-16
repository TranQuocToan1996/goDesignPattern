package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"

	ecies "github.com/ecies/go/v2"
)

type AsymEncryptionStrategy interface {
	Encrypt(data interface{}) (cipher []byte, err error)
}

type EllipticCurvestrategy struct{}
type RSA struct{}

// https://github.com/ecies/go
func (strat *EllipticCurvestrategy) Encrypt(data interface{}) (cipher []byte, err error) {
	msgBuf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return ecies.Encrypt(keyEllip.PublicKey, msgBuf)

}

// https://www.sohamkamani.com/golang/rsa-encryption/
func (strat *RSA) Encrypt(data interface{}) (cipher []byte, err error) {
	msgBuf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		msgBuf,
		nil,
	)
}

func EncryptMessage(plainText string, strategy AsymEncryptionStrategy) (cipher []byte, err error) {
	return strategy.Encrypt(plainText)
}
