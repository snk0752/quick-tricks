// Package license provides function to check if license file is exposed.
package license

import (
	"io"
	"net/http"
)

const endpoint = "/bitrix/license_key.php"

func Detect(target string) (string, error) {
	url := target + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, _ := io.ReadAll(resp.Body)
	if len(body) == 0 {
		return "", nil
	} else {
		return url, nil
	}
}
