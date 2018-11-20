package aliyun

import (
	"demo/lib"
	"demo/lib/uuid"
)

const accessKeyId string = "LTAIdN4r5tvSOKHE"
const accessKeySecret string = "t26Dr6xqZDOhBd0HvbXvl0ijfSEj5D"

var profile Profile = Profile{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret}
var Client IAliYunClient = DefaultClient{Profile: profile}

func InspectText(doc string, n int) []byte {
	slice1 := make([]string, 0)
	ll := doc
	for n < len([]rune(ll)) {
		rs := []rune(ll)
		m := lib.Substr2(ll, 0, n)
		slice1 = append(slice1, m)
		ll = lib.Substr2(string(rs), n, len(rs))
	}
	slice1 = append(slice1, ll)
	data := text(slice1)
	return data
}
func InspectImg(urls []string) []byte {
	return img(urls)
}

func text(docs []string) []byte {
	path := "/green/text/scan"
	clientInfo := ClinetInfo{Ip: "127.0.0.1"}
	bizType := "Green"
	scenes := []string{"antispam"}
	tasks := make([]TextTask, 0)
	for _, value := range docs {
		task := TextTask{
			DataId:   uuid.Rand().Hex(),
			Content:  value,
			Category: "post",
			Action:   "new",
		}
		tasks = append(tasks, task)

	}
	bizData := TextBizData{bizType, scenes, tasks}

	return Client.GetResponse(path, clientInfo, bizData)
}

func img(urls []string) []byte {
	path := "/green/image/scan"

	clientInfo := ClinetInfo{Ip: "127.0.0.1"}

	// 构造请求数据
	bizType := "Green"
	scenes := []string{"porn"}

	tasks := make([]ImgTask, 0)
	for _, value := range urls {
		task := ImgTask{
			DataId: uuid.Rand().Hex(),
			Url:    value,
		}
		tasks = append(tasks, task)
	}

	bizData := ImgBizData{bizType, scenes, tasks}
	return Client.GetResponse(path, clientInfo, bizData)
}
