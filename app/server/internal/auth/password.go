package auth

import (
	"fmt"
	"runtime"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	var hashParameters = argon2id.Params{
		Memory:      4 * 1024 * 1024, // 4 GiB of RAM
		Iterations:  8,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := argon2id.CreateHash(password, &hashParameters)
	if err != nil {
		return "", fmt.Errorf("Hashing error: %w", err)
	}

	return hash, nil
}

func CheckPassword(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}
