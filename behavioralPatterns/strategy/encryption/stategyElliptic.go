package encryption

import (
	ecies "github.com/ecies/go/v2"
)

var (
	// secp256k1 key pair
	keyEllip, _ = ecies.GenerateKey()
)
