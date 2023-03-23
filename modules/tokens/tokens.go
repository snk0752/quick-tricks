// Package tokens provides functions to get and parse Bitrix composite_data.php
package tokens

import (
	"fmt"
	"github.com/indigo-sadland/quick-tricks/utils/netclient"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const endpoint = "/bitrix/tools/composite_data.php"

type CompositeData struct {
	ServerTzOffset int
	ServerTime     int
	UserTzOffset   int
	BitrixSessid   string
}

func Get(target, proxy string) (*CompositeData, *http.Cookie, error) {
	var serverTzOffset, userTzOffset, serverTime int
	var bitrixSessid string
	var cookie *http.Cookie

	client, err := netclient.NewHTTPClient(proxy)
	if err != nil {
		err = fmt.Errorf("Unable to parse proxy string: %s", err.Error())
		return &CompositeData{}, &http.Cookie{}, err
	}
	url := target + endpoint
	resp, err := client.Get(url)
	if err != nil {
		return &CompositeData{}, &http.Cookie{}, err
	}

	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		raw := strings.Trim(string(body), "{")
		raw = strings.Trim(raw, "}")
		split := strings.Split(raw, ",")
		for _, v := range split {
			if strings.Contains(v, "SERVER_TZ_OFFSET") {
				secondSplit := strings.Split(v, ":")
				serverTzOffset, _ = strconv.Atoi(strings.ReplaceAll(secondSplit[1], "'", ""))
				continue
			} else if strings.Contains(v, "SERVER_TIME") {
				secondSplit := strings.Split(v, ":")
				serverTime, _ = strconv.Atoi(strings.ReplaceAll(secondSplit[1], "'", ""))
				continue
			} else if strings.Contains(v, "USER_TZ_OFFSET") {
				secondSplit := strings.Split(v, ":")
				userTzOffset, _ = strconv.Atoi(strings.ReplaceAll(secondSplit[1], "'", ""))
				continue
			} else if strings.Contains(v, "bitrix_sessid") {
				secondSplit := strings.Split(v, ":")
				bitrixSessid = strings.ReplaceAll(secondSplit[1], "'", "")
				continue
			} else {
				continue
			}
		}

		compositeData := &CompositeData{
			ServerTzOffset: serverTzOffset,
			ServerTime:     serverTime,
			UserTzOffset:   userTzOffset,
			BitrixSessid:   bitrixSessid,
		}
		cookies := resp.Cookies()
		for _, v := range cookies {
			if strings.Contains(v.Name, "PHPSESSID") {
				cookie = &http.Cookie{
					Name:  v.Name,
					Value: v.Value,
				}
			}
		}
		return compositeData, cookie, nil
	}

	return &CompositeData{}, &http.Cookie{}, nil
}
