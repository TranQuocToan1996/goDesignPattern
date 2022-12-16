package encryption

import (
	"crypto"
	"crypto/rsa"
	"log"
	"strings"
	"testing"

	ecies "github.com/ecies/go/v2"
)

func TestStrategy(t *testing.T) {
	// usage
	msg := "this is a confidential message"
	cipherRSA, err := EncryptMessage(msg, new(RSA))
	if err != nil {
		t.Fatal(err.Error())
	}

	// The first argument is an optional random data generator (the rand.Reader we used before)
	// we can set this value as nil
	// The OAEPOptions in the end signify that we encrypted the data using OAEP, and that we used
	// SHA256 to hash the input.
	decryptedBytes, err := privateKey.Decrypt(nil, cipherRSA, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		t.Fatal(err)
	}

	// We get back the original information in the form of bytes, which we
	// the cast to a string and print
	toStringRSA := string(decryptedBytes)
	if !strings.Contains(toStringRSA, msg) {
		t.Errorf("decryptedBytes - %v, msg - %v", string(decryptedBytes), msg)
	}

	cipherEllipticCurve, err := EncryptMessage(msg, new(EllipticCurvestrategy))
	if err != nil {
		t.Fatal(err.Error())
	}

	plaintextBuf, err := ecies.Decrypt(keyEllip, cipherEllipticCurve)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("ciphertext decrypted: %s\n", string(plaintextBuf))
	if !strings.Contains(string(plaintextBuf), msg) {
		t.Errorf("plaintextBuf - %v, msg - %v", string(plaintextBuf), msg)
	}
}
