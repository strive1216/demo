package aliyun

type Profile struct {
	AccessKeyId     string
	AccessKeySecret string
}

type ClinetInfo struct {
	SdkVersion  string `json:"sdkVersion,omitempty"`
	CfgVersion  string `json:"cfgVersion,omitempty"`
	UserType    string `json:"userType,omitempty"`
	UserId      string `json:"userId,omitempty"`
	UserNick    string `json:"userNick,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Imei        string `json:"imei,omitempty"`
	Imsi        string `json:"imsi,omitempty"`
	Umid        string `json:"umid,omitempty"`
	Ip          string `json:"ip,omitempty"`
	Os          string `json:"os,omitempty"`
	Channel     string `json:"channel,omitempty"`
	HostAppName string `json:"hostAppName,omitempty"`
	HostPackage string `json:"hostPackage,omitempty"`
	HostVersion string `json:"hostVersion,omitempty"`
}

type TextTask struct {
	DataId   string `json:"dataId,omitempty"`
	Content  string `json:"content,omitempty"`
	Category string `json:"category,omitempty"`
	Action   string `json:"action,omitempty"`
}

type ImgTask struct {
	DataId string `json:"dataId,omitempty"`
	Url    string `json:"url,omitempty"`
}

type TextBizData struct {
	BizType string     `json:"bizType,omitempty"`
	Scenes  []string   `json:"scenes,omitempty"`
	Tasks   []TextTask `json:"tasks,omitempty"`
}

type ImgBizData struct {
	BizType string    `json:"bizType,omitempty"`
	Scenes  []string  `json:"scenes,omitempty"`
	Tasks   []ImgTask `json:"tasks,omitempty"`
}

type TextResult struct {
	Code      int       `json:"code,omitempty"`
	Msg       string    `json:"msg,omitempty"`
	RequestId string    `json:"requestId,omitempty"`
	Data      []Textget `json:"data,omitempty"`
}

type Textget struct {
	Code    int          `json:"code,omitempty"`
	Msg     string       `json:"msg,omitempty"`
	DataId  string       `json:"dataId,omitempty"`
	TaskId  string       `json:"taskId,omitempty"`
	Content string       `json:"content,omitempty"`
	Results []textresult `json:"results,omitempty"`
}

type textresult struct {
	Scene      string      `json:"scene,omitempty"`
	Suggestion string      `json:"suggestion,omitempty"`
	Label      string      `json:"label,omitempty"`
	Rate       float32     `json:"rate,omitempty"`
	Extras     interface{} `json:"extras,omitempty"`
	Details    []detail    `json:"details,omitempty"`
}
type detail struct {
	Label    string    `json:"label,omitempty"`
	Contexts []context `json:"contexts,omitempty"`
}
type context struct {
	Context  string `json:"context,omitempty"`
	LibName  string `json:"libName,omitempty"`
	RuleType string `json:"ruleType,omitempty"`
}

type ImgResult struct {
	Code      int      `json:"code,omitempty"`
	Msg       string   `json:"msg,omitempty"`
	RequestId string   `json:"requestId,omitempty"`
	Data      []imgget `json:"data,omitempty"`
}
type imgget struct {
	Code    int               `json:"code,omitempty"`
	Msg     string            `json:"msg,omitempty"`
	DataId  string            `json:"dataId,omitempty"`
	TaskId  string            `json:"taskId,omitempty"`
	Url     string            `json:"url,omitempty"`
	Extras  map[string]string `json:"extras,omitempty"`
	Results []imgresult       `json:"results,omitempty"`
}

type imgresult struct {
	Scene      string      `json:"scene,omitempty"`
	Suggestion string      `json:"suggestion,omitempty"`
	Label      string      `json:"label,omitempty"`
	Rate       float32     `json:"rate,omitempty"`
	QrcodeData []string    `json:"qrcodeData,omitempty"`
	LogoData   []logData   `json:"logoData,omitempty"`
	SfaceData  []sfaceData `json:"sfaceData,omitempty"`
}

type logData struct {
	Type string  `json:"type,omitempty"`
	name string  `json:"name,omitempty"`
	X    float32 `json:"x,omitempty"`
	Y    float32 `json:"y,omitempty"`
	W    float32 `json:"w,omitempty"`
	H    float32 `json:"h,omitempty"`
}
type sfaceData struct {
	X     float32 `json:"x,omitempty"`
	Y     float32 `json:"y,omitempty"`
	W     float32 `json:"w,omitempty"`
	H     float32 `json:"h,omitempty"`
	Faces []faces `json:"faces,omitempty"`
}
type faces struct {
	Name string  `json:"name,omitempty"`
	Rate float32 `json:"rate,omitempty"`
	Id   string  `json:"id,omitempty"`
}
