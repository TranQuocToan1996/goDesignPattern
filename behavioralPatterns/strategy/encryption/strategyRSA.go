package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"io"
)

//TODO: Move to priv lib

func generationKeyRSA(reader ...io.Reader) (*rsa.PrivateKey, error) {
	var input io.Reader
	if len(reader) > 0 {
		input = io.MultiReader(reader...)
	} else {
		input = rand.Reader
	}

	// The GenerateKey method takes in a reader that returns random bits, and
	// the number of bitss
	return rsa.GenerateKey(input, 2048)
}

// The public key is a part of the *rsa.PrivateKey struct
var (
	privateKey, _ = generationKeyRSA()

	// The public key is a part of the *rsa.PrivateKey struct
	publicKey = privateKey.PublicKey
)
