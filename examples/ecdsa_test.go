package examples

/* Example code in GoLang for creating and verifying signatures using ECDSA
   See also https://blog.cloudflare.com/ecdsa-the-digital-signature-algorithm-of-a-better-internet/

	 Performance  (go 1.13)     sysbench      1000x Sign      1000x Verify
		Intel i5, 4570S             99971         39 ms             3 ms
		Pi-2 (no oc)                  222       6200 ms            41 ms
		Pi-2 (mild oc)                222				5400 ms            37 ms
		Pi-3B+ (no oc)      				  288 			7100 ms            45 ms
*/

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	_ "crypto/sha256" // https://github.com/Kong/go-srp/issues/1
	"encoding/asn1"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"
)

type ECDSASignature struct {
	R, S *big.Int
}

// Create a public/private key set
func createECDSAKeys() *ecdsa.PrivateKey {
	rng := rand.Reader
	curve := elliptic.P256()
	keys, _ := ecdsa.GenerateKey(curve, rng)
	return keys
}

// Create a message signature
func createECDSASignature(message []byte, priv *ecdsa.PrivateKey) ([]byte, error) {
	hashed := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hashed[:])
	if err != nil {
		return nil, err
	}
	return asn1.Marshal(ECDSASignature{r, s})
}

// Verify a signature
// See also https://leanpub.com/gocrypto/read#leanpub-auto-cryptographic-hashing-algorithms
func VerifyECDSASignature(message []byte, signature []byte, pub *ecdsa.PublicKey) bool {
	var rs ECDSASignature
	if _, err := asn1.Unmarshal(signature, &rs); err != nil {
		return false
	}

	hashed := sha256.Sum256(message)
	return ecdsa.Verify(pub, hashed[:], rs.R, rs.S)
}

// TestECDSASigning creates and verifies a signature
func TestECDSASigning(t *testing.T) {
	var keys = createECDSAKeys()
	var message = []byte("Hello World")
	var signature, _ = createECDSASignature(message, keys)
	log.Print("Created ECDSA message signature:", signature)

	success := VerifyECDSASignature(message, signature, &keys.PublicKey)
	log.Print("Verified ECDSA signature Result= ", success)
	if !success {
		t.Errorf("Signature ECDSA Verification Failed")
	}
}

// TestPerformance shows how many signatures can be made in 1 second
func TestECDSAPerformance(t *testing.T) {
	var keys = createECDSAKeys()
	var message = []byte("{ Hello World }")

	// signing
	start := time.Now()
	var sig []byte
	for i := 0; i < 1000; i++ {
		sig, _ = createECDSASignature(message, keys)
	}
	fmt.Println("Time to sign ECDSA 1000x:", time.Since(start))

	// verifying
	start = time.Now()
	var successCount int
	for i := 0; i < 1000; i++ {
		success := VerifyECDSASignature(sig, message, &keys.PublicKey)
		if success {
			successCount++
		}
	}
	fmt.Println("Time to verify ECDSA 1000x:", time.Since(start), successCount)
}

func main() {
	TestECDSAPerformance(nil)
}
