package examples

/* Example code in GoLang for creating and verifying signatures
 See also:
-   [https://golang.org/pkg/crypto/rsa/\#GenerateKey](https://golang.org/pkg/crypto/rsa/#GenerateKey)

-   [https://golang.org/pkg/crypto/rsa/\#SignPSS](https://golang.org/pkg/crypto/rsa/#SignPSS)

-   [https://golang.org/pkg/crypto/rsa/\#VerifyPSS](https://golang.org/pkg/crypto/rsa/#VerifyPSS)


Use the golang crypto library: https://golang.org/src/crypto
* https://github.com/brainattica/Golang-RSA-sample/blob/master/rsa_sample.go

*/

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

// Create a message signature:
func createSignature() {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	var newHash = crypto.SHA256
	pssHash := newhash.New()
	pssHash.Write(message)
	messageHash := pssHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, newhash, messageHash, &opts)
}

// Verify a signature
func verifySignature() {
	var publisherNode = myzone.getPublisher(payload.message.sender)
	var publicKey = publisherNode.publicKey // from discovery
	var signature = payload.signature

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	var newHash = crypto.SHA256
	pssHash := newhash.New()
	pssHash.Write(message)
	messageHash := pssHash.Sum(nil)

	var err = rsa.VerifyPSS(publicKey, crypto.SHA256, messageHash, signature, &opts)
}
