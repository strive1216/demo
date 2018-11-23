package aliyun

import (
	"demo/lib/uuid"
)

const accessKeyId string = "LTAIdN4r5tvSOKHE"
const accessKeySecret string = "t26Dr6xqZDOhBd0HvbXvl0ijfSEj5D"

var profile Profile = Profile{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret}
var Client IAliYunClient = DefaultClient{Profile: profile}

func InspectText(doc []string, bizType string, scenes []string) []byte {
	data := text(doc, bizType, scenes)
	return data
}
func InspectImg(urls []string, bizType string, scenes []string) []byte {
	return img(urls, bizType, scenes)
}

func text(docs []string, bizType string, scenes []string) []byte {
	path := "/green/text/scan"
	clientInfo := ClinetInfo{Ip: "127.0.0.1"}
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

func img(urls []string, bizType string, scenes []string) []byte {
	path := "/green/image/scan"

	clientInfo := ClinetInfo{Ip: "127.0.0.1"}

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
