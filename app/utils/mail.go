package utils

import (
	"io"
	"net/http"
)

func DownloadMailTemplate(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()

	cont, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}

	return cont, nil
}
