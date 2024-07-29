package helper

import (
	"crypto/rand"
	"io"
	"math/big"
	"strings"

	"golang.org/x/crypto/chacha20"
)

// Got this shit from Claude, idk how it operates but it encrypts array of bytes via ChaCha20 encryption with a specefic key
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	// Generate a random nonce
	nonce := make([]byte, chacha20.NonceSizeX)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Create the ChaCha20 cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)

	// Prepend the nonce to the ciphertext
	return append(nonce, ciphertext...), nil
}

// Got this shit from Claude, idk how it operates but it decrypts array of bytes via ChaCha20 encryption with a specefic key
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	// Extract the nonce from the ciphertext
	nonce, ciphertext := ciphertext[:chacha20.NonceSizeX], ciphertext[chacha20.NonceSizeX:]

	// Create the ChaCha20 cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext
	plaintext := make([]byte, len(ciphertext))
	cipher.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

// Got from clause, generates 32-length password
func GeneratePassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?"

	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	// Ensure the password contains at least one of each character type
	types := []string{"[a-z]", "[A-Z]", "[0-9]", "[!@#$%^&*()-_=+\\[\\]{}|;:,.<>?]"}
	for _, t := range types {
		if !strings.ContainsAny(string(password), t) {
			randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
			charOfType, _ := rand.Int(rand.Reader, big.NewInt(int64(len(t)-2))) // -2 for the brackets
			password[randomIndex.Int64()] = t[charOfType.Int64()+1]             // +1 to skip the opening bracket
		}
	}

	return string(password), nil
}
