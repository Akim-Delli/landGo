package contentType

import (
	"fmt"
	"net/http"
)

func ContentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")

	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type Header")
	}

	return resp.Header.Get("Content-Type"), nil

}
