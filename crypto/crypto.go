package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
)

func EncryptFile(content string, password string, outputDir string, fileName string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return err
	}

	key := deriveKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(content), nil)

	finalData := append(salt, ciphertext...)

	encodedData := base64.StdEncoding.EncodeToString(finalData)

	outputPath := filepath.Join(outputDir, fileName+".txt")
	return os.WriteFile(outputPath, []byte(encodedData), 0644)
}

func DecryptFile(filePath string, password string) (string, error) {
	encodedData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Decode base64
	data, err := base64.StdEncoding.DecodeString(string(encodedData))
	if err != nil {
		return "", err
	}

	// Extract salt and ciphertext
	if len(data) < 16 {
		return "", errors.New("invalid encrypted data")
	}
	salt := data[:16]
	ciphertext := data[16:]

	// Derive key from password
	key := deriveKey(password, salt)

	// Create cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract nonce
	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("invalid ciphertext")
	}
	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// deriveKey derives a 32-byte key from the password and salt
func deriveKey(password string, salt []byte) []byte {
	// In a real application, you should use a proper key derivation function like PBKDF2
	// For this example, we'll use a simple XOR-based derivation
	key := make([]byte, 32)
	passwordBytes := []byte(password)
	
	for i := 0; i < 32; i++ {
		key[i] = passwordBytes[i%len(passwordBytes)] ^ salt[i%len(salt)]
	}
	
	return key
} 