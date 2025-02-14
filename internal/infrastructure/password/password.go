package password

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	std_errors "errors"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"golang.org/x/crypto/argon2"
)

type hasher struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
	saltLength  uint32
}

func New() *hasher {
	return &hasher{
		memory:      16 * 1024,
		iterations:  3,
		parallelism: 8,
		keyLength:   32,
		saltLength:  16,
	}
}

func (h *hasher) generateSalt() ([]byte, error) {
	salt := make([]byte, h.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func (h *hasher) HashPassword(password string) (string, error) {
	salt, err := h.generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, h.iterations, h.memory, h.parallelism, h.keyLength)

	return encodeHash(salt, hash), nil
}

func (h *hasher) ComparePassword(hashedPassword, password string) error {
	salt, expectedHash, err := decodeHash(hashedPassword)
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(password), salt, h.iterations, h.memory, h.parallelism, h.keyLength)

	if !bytes.Equal(hash, expectedHash) {
		return errors.ErrInvalidPassword
	}

	return nil
}

func encodeHash(salt, hash []byte) string {
	return base64.StdEncoding.EncodeToString(salt) + "$" + base64.StdEncoding.EncodeToString(hash)
}

func decodeHash(encoded string) ([]byte, []byte, error) {
	parts := bytes.Split([]byte(encoded), []byte("$"))
	if len(parts) != 2 {
		return nil, nil, std_errors.New("invalid hash format")
	}

	salt, err := base64.StdEncoding.DecodeString(string(parts[0]))
	if err != nil {
		return nil, nil, err
	}

	hash, err := base64.StdEncoding.DecodeString(string(parts[1]))
	if err != nil {
		return nil, nil, err
	}

	return salt, hash, nil
}
