package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func sha256Hashing(input string) string {
	plainText := []byte(input)
	sha256Sum := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Sum[:])
}

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func encryptIt(value []byte, key string) []byte {
	aesBlock, err := aes.NewCipher([]byte(mdHashing(key)))

	if err != nil {
		fmt.Printf("Err generating aesBlock: %s\n", err)
		return nil
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)

	if err != nil {
		fmt.Printf("Error generating gcmInstance: %s\n", err)
		return nil
	}

	//generate random number into the nonce
	nonce := make([]byte, gcmInstance.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)
	return cipheredText
}

func main() {
	fmt.Println(sha256Hashing("Hello_world"))
	fmt.Println(sha256Hashing("Silly_me"))

	fmt.Println("---------------Bytes---------------")
	fmt.Println(encryptIt([]byte("This is some random text"), "random.key"))
	fmt.Println("------------------------------------")

	fmt.Println("---------------String---------------")
	fmt.Println(string(encryptIt([]byte("This is some random text"), "random.key")))
	fmt.Println("---------------------------------------")
}
