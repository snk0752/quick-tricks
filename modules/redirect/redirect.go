package redirect

import (
	"io"
	"net/http"
	"strings"
)
const (
	endpoint1 = "/bitrix/redirect.php?goto="
	endpoint2 = "/bitrix/rk.php?goto="
	endpoint3 = "/bitrix/tools/track_mail_click.php?url="
)
func Detect(target string) ([]string, error) {
	var pages []string
	endpoints := []string{endpoint1, endpoint2, endpoint3}

	for _, v := range endpoints {
		url := target + v + target + "%252F:123@attacker.example"
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		} else {
			if resp.StatusCode == 404 {
				body, _ := io.ReadAll(resp.Body)
				if len(body) != 0 {
					if strings.Contains(string(body), "Внимание! Вы перенаправляетесь на другой сайт") {
						pages = append(pages, url)
					}
				}
			}
		}

	}
	return pages, nil
}