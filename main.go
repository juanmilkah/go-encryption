package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func sha256Hashing(input string) string {
	plainText := []byte(input)
	sha256Sum := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Sum[:])
}

func main() {
	fmt.Println(sha256Hashing("Hello_world"))
	fmt.Println(sha256Hashing("Silly_me"))
}
