package aliyun

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"demo/lib/uuid"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const host string = "http://green.cn-shanghai.aliyuncs.com"
const method string = "POST"
const newline string = "\n"
const MIME string = "application/json"

func addRequestHeader(requestBody string, req *http.Request, clientInfo string, path string, accessKeyId string, accessKeySecret string) {
	var now = time.Now().UTC()
	var gmtDate = string([]byte(now.Weekday().String())[0:3]) + now.Format(", 02 Jan 2006 15:04:05 GMT")
	var md5Ctx = md5.New()
	md5Ctx.Write([]byte(requestBody))
	cipherStr := md5Ctx.Sum(nil)
	base64Md5Str := base64.StdEncoding.EncodeToString(cipherStr)

	acsHeaderKeyArray := []string{"x-acs-signature-method", "x-acs-signature-nonce", "x-acs-signature-version", "x-acs-version"}
	acsHeaderValueArray := []string{"HMAC-SHA1", uuid.Rand().Hex(), "1.0", "2018-05-09"}

	req.Header.Set("Accept", MIME)
	req.Header.Set("Content-Type", MIME)
	req.Header.Set("Content-Md5", base64Md5Str)
	req.Header.Set("Date", gmtDate)
	req.Header.Set(acsHeaderKeyArray[0], acsHeaderValueArray[0])
	req.Header.Set(acsHeaderKeyArray[1], acsHeaderValueArray[1])
	req.Header.Set(acsHeaderKeyArray[2], acsHeaderValueArray[2])
	req.Header.Set(acsHeaderKeyArray[3], acsHeaderValueArray[3])
	req.Header.Set("Authorization", "acs"+" "+accessKeyId+":"+singature(acsHeaderKeyArray, acsHeaderValueArray, base64Md5Str, clientInfo, path, accessKeySecret, gmtDate))
}

func singature(acsHeaderKeyArray []string, acsHeaderValueArray []string, md5Str string, clientInfo string, path string, accessKeySecret string, gmtDate string) string {
	b := bytes.Buffer{}

	b.WriteString(method)
	b.WriteString(newline)

	b.WriteString(MIME)
	b.WriteString(newline)

	b.WriteString(md5Str)
	b.WriteString(newline)

	b.WriteString(MIME)
	b.WriteString(newline)

	b.WriteString(gmtDate)
	b.WriteString(newline)

	for i := 0; i < len(acsHeaderKeyArray); i++ {
		b.WriteString(acsHeaderKeyArray[i])
		b.WriteString(":")
		b.WriteString(acsHeaderValueArray[i])
		b.WriteString(newline)
	}

	b.WriteString(path)
	b.WriteString("?clientInfo=")
	b.WriteString(clientInfo)

	mac := hmac.New(sha1.New, []byte(accessKeySecret))
	mac.Write([]byte(b.String()))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func ErrorResult(error error) string {
	errorResult := make(map[string]string)
	errorResult["code"] = "500"
	errorResult["msg"] = error.Error()
	errorJson, _ := json.Marshal(errorResult)
	return string(errorJson)
}

type DefaultClient struct {
	Profile Profile
}

func (defaultClient DefaultClient) GetResponse(path string, clinetInfo interface{}, bizData interface{}) []byte {

	clientInfoJson, _ := json.Marshal(clinetInfo)
	bizDataJson, _ := json.Marshal(bizData)

	client := &http.Client{}
	req, err := http.NewRequest(method, host+path+"?clientInfo="+url.QueryEscape(string(clientInfoJson)), strings.NewReader(string(bizDataJson)))

	if err != nil {
		fmt.Println(err.Error())
		return []byte(ErrorResult(err))
	} else {
		addRequestHeader(string(bizDataJson), req, string(clientInfoJson), path, accessKeyId, accessKeySecret)

		response, _ := client.Do(req)

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			// handle error
			fmt.Println(err.Error())
			return []byte(ErrorResult(err))
		} else {
			return body
		}
	}
}

type IAliYunClient interface {
	GetResponse(path string, clinetInfo interface{}, bizData interface{}) []byte
}
