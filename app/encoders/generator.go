package encoders

import "crypto/rand"

func GenerateStandardCode(length int) (string, error) {

	chars := "0123456789"

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, j := range bytes {
		bytes[i] = chars[j%byte(len(chars))]
	}
	return string(bytes), nil
}

func GenerateAlphanumericCode(length int) (string, error) {

	chars := "ABCDEFJHIJKLMNOPQRSTUVWXYZabcedfghijklmnopqrstuvwxyz0123456789"

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, j := range bytes {
		bytes[i] = chars[j%byte(len(chars))]
	}
	return string(bytes), nil
}
