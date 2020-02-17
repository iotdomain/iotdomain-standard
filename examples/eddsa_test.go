package examples

/* Example code in GoLang for creating and verifying signatures using EdDSA (Ed25519)

https://pkg.go.dev/github.com/katzenpost/core/crypto/eddsa?tab=doc, and
https://pkg.go.dev/golang.org/x/crypto/ed25519?tab=doc

 Performance (go 1.13)              1000x Sign      1000x Verify
	Intel i5, 4570S (5800 BogoMIPS)     62 ms             0.16 ms
	Pi-2  (mild oc)                   2000 ms          6300 ms
	Pi-3  (38 BogoMips)               2800 ms          8800 ms
*/

import (
	"crypto/ed25519"
	"crypto/rand"
	_ "crypto/sha256" // https://github.com/Kong/go-srp/issues/1
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"
)

type Ed25519Signature struct {
	R, S *big.Int
}

// Create a public/private key set
func createEd25519Keys() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	rng := rand.Reader
	pubKey, privKey, err := ed25519.GenerateKey(rng)
	return pubKey, privKey, err
}

// Create a message signature
func createEd25519Signature(message []byte, priv ed25519.PrivateKey) []byte {
	sig := ed25519.Sign(priv, message)
	return sig
}

// Verify a signature
// See also https://leanpub.com/gocrypto/read#leanpub-auto-cryptographic-hashing-algorithms
func VerifyEd25519Signature(message []byte, signature []byte, pub ed25519.PublicKey) bool {
	return ed25519.Verify(pub, message, signature)
}

// TestEdDSASigning creates and verifies a signature
func TestEd25519Signing(t *testing.T) {
	var pubKey, privKey, _ = createEd25519Keys()
	var message = []byte("Hello World")
	var signature = createEd25519Signature(message, privKey)
	log.Print("Created Ed25519 message signature:", signature)

	success := VerifyEd25519Signature(message, signature, pubKey)
	log.Print("Verified Ed25519 signature Result= ", success)
	if !success {
		t.Errorf("Signature Ed25519 Verification Failed")
	}
}

// TestPerformance shows how many signatures can be made in 1 second
func TestEd25519Performance(t *testing.T) {
	var pubKey, privKey, _ = createEd25519Keys()
	var message = []byte("{ Hello World }")

	// signing
	start := time.Now()
	var signature []byte
	for i := 0; i < 1000; i++ {
		signature = createEd25519Signature(message, privKey)
	}
	fmt.Println("Time to sign Ed25519 1000x:", time.Since(start))

	// verifying
	start = time.Now()
	for i := 0; i < 1000; i++ {
		VerifyEd25519Signature(message, signature, pubKey)
	}
	fmt.Println("Time to verify Ed25519 1000x:", time.Since(start))
}
