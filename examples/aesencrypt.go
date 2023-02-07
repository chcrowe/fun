package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {

	// cipher key
	key := "thisis32bitlongpassphraseimusing"
	// plaintext
	fmt.Printf("%-12s%s\n", "KEY:", key)

	// plaintext
	pt := "This is a secret"

	c := EncryptAES([]byte(key), pt)

	// plaintext
	fmt.Printf("%-12s%s\n", "PLAIN:", pt)

	// ciphertext
	fmt.Printf("%-12s%s\n", "CIPHER:", c)

	// decrypt
	dec := DecryptAES([]byte(key), c)

	fmt.Printf("%-12s%s\n", "DECRYPTED:", dec)

}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
