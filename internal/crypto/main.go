package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/scrypt"
)

const (
	magic       = "BUNKR\x00"
	saltLen     = 16
	nonceLen    = 12
	derivedKeyL = 32
)

var (
	scryptN = 1 << 15
	scryptR = 8
	scryptP = 1
)

func DeriveKey(passphrase string, salt []byte) ([]byte, error) {
	if len(salt) != saltLen {
		return nil, fmt.Errorf("salt must be %d bytes", saltLen)
	}
	return scrypt.Key([]byte(passphrase), salt, scryptN, scryptR, scryptP, derivedKeyL)
}

func EncryptBytes(passphrase string, plaintext []byte) ([]byte, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	key, err := DeriveKey(passphrase, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, nonceLen)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	out := append([]byte(magic), salt...)
	out = append(out, nonce...)
	out = append(out, ciphertext...)
	return out, nil
}

func DecryptBytes(passphrase string, blob []byte) ([]byte, error) {
	minLen := len(magic) + saltLen + nonceLen
	if len(blob) < minLen {
		return nil, errors.New("ciphertext too short or corrupt")
	}

	if string(blob[:len(magic)]) != magic {
		return nil, errors.New("invalid file format")
	}

	offset := len(magic)
	salt := blob[offset : offset+saltLen]
	offset += saltLen
	nonce := blob[offset : offset+nonceLen]
	offset += nonceLen
	ciphertext := blob[offset:]

	key, err := DeriveKey(passphrase, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("decryption failed: wrong password or corrupted file")
	}
	return plaintext, nil
}

func EncryptFile(passphrase, inPath, outPath string) error {
	data, err := os.ReadFile(inPath)
	if err != nil {
		return err
	}
	blob, err := EncryptBytes(passphrase, data)
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, blob, 0600)
}

func DecryptFile(passphrase, inPath, outPath string) error {
	blob, err := os.ReadFile(inPath)
	if err != nil {
		return err
	}
	data, err := DecryptBytes(passphrase, blob)
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, data, 0600)
}
