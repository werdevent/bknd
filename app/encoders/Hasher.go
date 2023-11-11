package encoders

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {

	cost := 8

	hashPWD, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", err
	}

	return string(hashPWD), nil
}

func CompareHashedPWD(hashed, pwd string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))
	if err != nil {

		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, errors.New("incorrect password")

		case errors.Is(err, bcrypt.ErrHashTooShort):
			return false, errors.New("the password introduced is invalid")
		default:
			return false, err
		}

	}

	return true, nil
}
