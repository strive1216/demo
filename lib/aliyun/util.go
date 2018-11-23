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

type DefaultClient struct {
	Profile Profile
}

const host string = "http://green.cn-shanghai.aliyuncs.com"
const method string = "POST"
const newline string = "\n"
const MIME string = "application/json"

func ErrorResult(error error) string {
	errorResult := make(map[string]string)
	errorResult["code"] = "500"
	errorResult["msg"] = error.Error()
	errorJson, _ := json.Marshal(errorResult)
	return string(errorJson)
}

func (defaultClient DefaultClient) GetResponse(path string, clinetInfo interface{}, bizData interface{}) []byte {
	clientInfoJson, _ := json.Marshal(clinetInfo)
	bizDataJson, _ := json.Marshal(bizData)

	client := &http.Client{
		Timeout: time.Duration(20 * time.Second),
	}
	req, err := http.NewRequest(method, host+path+"?clientInfo="+url.QueryEscape(string(clientInfoJson)), strings.NewReader(string(bizDataJson)))

	if err != nil {
		fmt.Println(err.Error())
		return []byte(ErrorResult(err))
	} else {
		//addRequestHeader(string(bizDataJson), req, string(clientInfoJson), path, accessKeyId, accessKeySecret)
		HeaderMap := getheadermap(string(bizDataJson), string(clientInfoJson), path, accessKeyId, accessKeySecret)
		for k, v := range HeaderMap {
			req.Header.Set(k, v)
		}
		response, err_ := client.Do(req)
		if err_ != nil {
			return []byte(ErrorResult(err_))
		}
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

func (defaultClient DefaultClient) GetHeaderMap(path string, clinetInfo interface{}, bizData interface{}) map[string]string {
	clientInfoJson, _ := json.Marshal(clinetInfo)
	bizDataJson, _ := json.Marshal(bizData)

	return getheadermap(string(bizDataJson), string(clientInfoJson), path, accessKeyId, accessKeySecret)
}

/*获取阿里云内容安全公共参数*/
func getheadermap(requestBody string, clientInfo string, path string, accessKeyId string, accessKeySecret string) map[string]string {
	now := time.Now().UTC()
	gmtDate := string([]byte(now.Weekday().String())[0:3]) + now.Format(", 02 Jan 2006 15:04:05 GMT")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(requestBody))
	cipherStr := md5Ctx.Sum(nil)
	base64Md5Str := base64.StdEncoding.EncodeToString(cipherStr)

	//按字典排序 方便后续签名
	acsHeaderKeyArray := []string{"x-acs-signature-method", "x-acs-signature-nonce", "x-acs-signature-version", "x-acs-version"}
	acsHeaderValueArray := []string{"HMAC-SHA1", uuid.Rand().Hex(), "1.0", "2018-05-09"}

	HeaderMap := make(map[string]string)
	HeaderMap["Accept"] = MIME
	HeaderMap["Content-Type"] = MIME
	HeaderMap["Content-Md5"] = base64Md5Str
	HeaderMap["Date"] = gmtDate
	HeaderMap[acsHeaderKeyArray[0]] = acsHeaderValueArray[0]
	HeaderMap[acsHeaderKeyArray[1]] = acsHeaderValueArray[1]
	HeaderMap[acsHeaderKeyArray[2]] = acsHeaderValueArray[2]
	HeaderMap[acsHeaderKeyArray[3]] = acsHeaderValueArray[3]
	HeaderMap["Authorization"] = "acs " + accessKeyId + ":" + singature(acsHeaderKeyArray, acsHeaderValueArray, base64Md5Str, clientInfo, path, accessKeySecret, gmtDate)
	return HeaderMap
}

/*获取阿里云内容安全签名*/
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
