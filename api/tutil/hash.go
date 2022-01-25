package tutil

import (
	"crypto/rand"
	"errors"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func HashMatches(password string, hashedPassword []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789" + 
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + 
		"abcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
