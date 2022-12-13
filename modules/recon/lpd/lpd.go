// Package lpd provides function to check urls where local path disclosure appears.
package lpd

import (
	"io"
	"net/http"
	"strings"
)

const (
	endpoint1 = "/?USER_FIELD_MANAGER=1"
	endpoint2 = "/bitrix/admin/restore_export.php"
	endpoint3 = "/bitrix/admin/tools_index.php"
	endpoint4 = "/bitrix/bitrix.php"
	endpoint5 = "/bitrix/modules/main/ajax_tools.php"
	endpoint6 = "/bitrix/php_interface/after_connect_d7.php"
	endpoint7 = "/bitrix/themes/.default/.description.php"
	endpoint8 = "/bitrix/components/bitrix/main.ui.selector/templates/.default/template.php"
	endpoint9 = "/bitrix/components/bitrix/forum.user.profile.edit/templates/.default/interface.php"
)

func Detect(target string) ([]string, error) {
	var pages []string
	endpoints := []string{endpoint1, endpoint2, endpoint3, endpoint4, endpoint5, endpoint6, endpoint7, endpoint8,
		endpoint9}

	for _, v := range endpoints {
		url := target + v
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		} else {
			if resp.StatusCode == 500 {
				body, _ := io.ReadAll(resp.Body)
				if len(body) != 0 {
					if !strings.Contains(string(body), "The script encountered an error and will be aborted") {
						pages = append(pages, url)
					}
				}
			} else {
				continue
			}
		}
	}

	return pages, nil
}
