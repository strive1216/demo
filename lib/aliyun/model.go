package aliyun

type Profile struct {
	AccessKeyId     string
	AccessKeySecret string
}

type ClinetInfo struct {
	SdkVersion  string `json:"sdkVersion"`
	CfgVersion  string `json:"cfgVersion"`
	UserType    string `json:"userType"`
	UserId      string `json:"userId"`
	UserNick    string `json:"userNick"`
	Avatar      string `json:"avatar"`
	Imei        string `json:"imei"`
	Imsi        string `json:"imsi"`
	Umid        string `json:"umid"`
	Ip          string `json:"ip"`
	Os          string `json:"os"`
	Channel     string `json:"channel"`
	HostAppName string `json:"hostAppName"`
	HostPackage string `json:"hostPackage"`
	HostVersion string `json:"hostVersion"`
}

type TextTask struct {
	DataId   string `json:"dataId"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Action   string `json:"action"`
}

type ImgTask struct {
	DataId string `json:"dataId"`
	Url    string `json:"url"`
}

type TextBizData struct {
	BizType string     `json:"bizType"`
	Scenes  []string   `json:"scenes"`
	Tasks   []TextTask `json:"tasks"`
}

type ImgBizData struct {
	BizType string    `json:"bizType"`
	Scenes  []string  `json:"scenes"`
	Tasks   []ImgTask `json:"tasks"`
}

type TextResult struct {
	Code      int       `json:"code"`
	Msg       string    `json:"msg"`
	RequestId string    `json:"requestId"`
	Data      []Textget `json:"data"`
}

type Textget struct {
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	DataId  string       `json:"dataId"`
	TaskId  string       `json:"taskId"`
	Content string       `json:"content"`
	Results []textresult `json:"results"`
}

type textresult struct {
	Scene      string      `json:"scene"`
	Suggestion string      `json:"suggestion"`
	Label      string      `json:"label"`
	Rate       float32     `json:"rate"`
	Extras     interface{} `json:"extras"`
	Details    []detail    `json:"details"`
}

type detail struct {
	Label    string    `json:"label"`
	Contexts []context `json:"contexts"`
}
type context struct {
	Context  string `json:"context"`
	LibName  string `json:"libName"`
	RuleType string `json:"ruleType"`
}
