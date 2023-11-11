package generators

import (
	"crypto/rand"
	"strconv"
)

func GenerateSixDigitsCode() (int32, error) {
	lenght := 6

	chars := "1234567890"

	buffer := make([]byte, lenght)

	_, err := rand.Read(buffer)
	if err != nil {
		return int32(0), err
	}

	for i, v := range buffer {
		buffer[i] = chars[v%byte(len(chars))]
	}

	code, _ := strconv.Atoi(string(buffer))

	return int32(code), nil
}
