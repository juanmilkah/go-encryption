# Cryptographic Utilities

A Go package that provides basic cryptographic operations including hashing (SHA-256 and MD5) and encryption/decryption using AES-GCM.

## Features

- SHA-256 hashing
- MD5 hashing (Note: MD5 is not cryptographically secure, used here for key derivation only)
- AES-GCM encryption and decryption
- Secure random nonce generation

## Installation

```go
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
```

## Usage

### Hashing Functions

#### SHA-256 Hashing

```go
func sha256Hashing(input string) string {
    plainText := []byte(input)
    sha256Sum := sha256.Sum256(plainText)
    return hex.EncodeToString(sha256Sum[:])
}
```

Example:
```go
hash := sha256Hashing("Hello_world")
fmt.Println(hash)
```

#### MD5 Hashing

```go
func mdHashing(input string) string {
    byteInput := []byte(input)
    md5Hash := md5.Sum(byteInput)
    return hex.EncodeToString(md5Hash[:])
}
```

### Encryption/Decryption Functions

#### Encryption

```go
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
    nonce := make([]byte, gcmInstance.NonceSize())
    io.ReadFull(rand.Reader, nonce)
    cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)
    return cipheredText
}
```

#### Decryption

```go
func decryptIt(cipherText []byte, key string) []byte {
    hashedKey := mdHashing(key)
    aesBlock, err := aes.NewCipher([]byte(hashedKey))
    if err != nil {
        fmt.Printf("Error generating aesBlock, decrypt: %s\n", err)
        return nil
    }
    gcmInstance, err := cipher.NewGCM(aesBlock)
    if err != nil {
        fmt.Printf("Error generating gcmInstance, decrypt: %s\n", err)
        return nil
    }
    nonceSize := gcmInstance.NonceSize()
    nonce, cipheredText := cipherText[:nonceSize], cipherText[nonceSize:]
    originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
    if err != nil {
        fmt.Printf("Error opening gcmInstance: %s\n", err)
        return nil
    }
    return originalText
}
```

### Complete Example

```go
func main() {
    // Hashing examples
    fmt.Println(sha256Hashing("Hello_world"))
    fmt.Println(sha256Hashing("Silly_me"))

    // Encryption/Decryption example
    key := "random.key"
    encryptedText := encryptIt([]byte("This is some random text"), key)
    decryptedText := decryptIt(encryptedText, key)

    // Print results
    fmt.Println("-------------Encrypted----------------")
    fmt.Println("---------------Bytes---------------")
    fmt.Println(encryptedText)
    fmt.Println("---------------String---------------")
    fmt.Println(string(encryptedText))
    fmt.Println("----------------Decrypted---------------")
    fmt.Println("----------------Bytes--------------------")
    fmt.Println(decryptedText)
    fmt.Println("----------------String-------------------")
    fmt.Println(string(decryptedText))
}
```

## Security Considerations

1. This implementation uses AES-GCM (Galois/Counter Mode) which provides both confidentiality and authenticity.
2. The nonce (number used once) is generated using a cryptographically secure random number generator.
3. MD5 is used only for key derivation and not for security-critical hashing operations.
4. The SHA-256 implementation is suitable for general cryptographic hashing needs.

## Error Handling

The code includes basic error handling for:
- AES cipher block creation
- GCM instance creation
- Decryption operations

All errors are logged to stdout using `fmt.Printf`.

## Limitations

1. The key length must be compatible with AES (will be hashed to appropriate length).
2. No key stretching is implemented - consider using a proper key derivation function for production use.
3. The encrypted output includes the nonce prepended to the ciphertext.

## Contributing

Feel free to submit issues and enhancement requests.

## License

[LICENSE](LICENSE)
