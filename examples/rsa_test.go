package examples

/* Example code in GoLang for creating and verifying signatures using RSA-PSS
* https://github.com/brainattica/Golang-RSA-sample/blob/master/rsa_sample.go

  Performance of signing is very poor:
  On Quad Core i5,4570S:
   	Time to sign 100x: 2100ms
		Time to verify 100x: 80ms

	On Raspberry Pi 3:
		tbd
*/

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256" // https://github.com/Kong/go-srp/issues/1
	"fmt"
	"log"
	"testing"
	"time"
)

// Create a public/private key set
// See also: https://golang.org/pkg/crypto/rsa/\#GenerateKey](https://golang.org/pkg/crypto/rsa/#GenerateKey
func createRSAKeys() *rsa.PrivateKey {
	rng := rand.Reader
	keys, _ := rsa.GenerateKey(rng, 2048)
	log.Print("Created public/private key pair")
	return keys
}

// Create a message signature
// See also: https://golang.org/pkg/crypto/rsa/\#SignPSS](https://golang.org/pkg/crypto/rsa/#SignPSS
func createRSASignature(message []byte, privateKey *rsa.PrivateKey) []byte {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	var newHash = crypto.SHA256
	pssHash := newHash.New()
	pssHash.Write(message)
	messageHash := pssHash.Sum(nil)

	signature, _ := rsa.SignPSS(rand.Reader, privateKey, newHash, messageHash, &opts)
	return signature
}

// Verify a signature
// See also: https://golang.org/pkg/crypto/rsa/\#VerifyPSS](https://golang.org/pkg/crypto/rsa/#VerifyPSS
func verifyRSASignature(signature []byte, message []byte, publicKey *rsa.PublicKey) bool {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	var newHash = crypto.SHA256
	pssHash := newHash.New()
	pssHash.Write(message)
	messageHash := pssHash.Sum(nil)

	var err = rsa.VerifyPSS(publicKey, crypto.SHA256, messageHash, signature, &opts)
	return (err == nil)
}

// TestRSASigning creates and verifies a RSA signature
func TestRSASigning(t *testing.T) {
	var keys = createRSAKeys()
	var message = []byte("Hello World")
	var signature = createRSASignature(message, keys)
	log.Print("Created RSA message signature:", signature)

	success := verifyRSASignature(signature, message, &keys.PublicKey)
	log.Print("Verified RSA signature Result= ", success)
	if !success {
		t.Errorf("Signature RSA Verification Failed")
	}
}

// TestPerformance shows how many signatures can be made in 1 second
func TestRSAPerformance(t *testing.T) {
	var keys = createRSAKeys()
	var message = []byte("{ Hello World }")

	// signing
	start := time.Now()
	var sig []byte
	for i := 0; i < 1000; i++ {
		sig = createRSASignature(message, keys)
	}
	fmt.Println("Time to sign RSA 1000x:", time.Since(start))

	// verifying
	start = time.Now()
	for i := 0; i < 1000; i++ {
		verifyRSASignature(sig, message, &keys.PublicKey)
	}
	fmt.Println("Time to verify RSA 1000x:", time.Since(start))
}
