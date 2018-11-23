package aliyun

type IAliYunClient interface {
	GetResponse(path string, clinetInfo interface{}, bizData interface{}) []byte
	GetHeaderMap(path string, clientInfo interface{}, bizData interface{}) map[string]string
}
