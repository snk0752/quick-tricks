// Package license provides function to check if license file is exposed.
package license

import (
	"fmt"
	"github.com/indigo-sadland/quick-tricks/utils/netclient"
	"io"
)

const endpoint = "/bitrix/license_key.php"

func Detect(target, proxy string) (string, error) {
	url := target + endpoint

	client, err := netclient.NewHTTPClient(proxy)
	if err != nil {
		err = fmt.Errorf("Unable to parse proxy string: %s", err.Error())
		return "", err
	}
	resp, err := client.Get(url)

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
