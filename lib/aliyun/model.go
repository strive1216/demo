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

type TextBizData struct {
	BizType string     `json:"bizType,omitempty"`
	Scenes  []string   `json:"scenes,omitempty"`
	Tasks   []TextTask `json:"tasks,omitempty"`
}
type TextTask struct {
	DataId   string `json:"dataId,omitempty"`
	Content  string `json:"content,omitempty"`
	Category string `json:"category,omitempty"`
	Action   string `json:"action,omitempty"`
}

type TextResult struct {
	Code      int        `json:"code,omitempty"`
	Msg       string     `json:"msg,omitempty"`
	RequestId string     `json:"requestId,omitempty"`
	Data      []textdata `json:"data,omitempty"`
}
type textdata struct {
	Code    int          `json:"code,omitempty"`
	Msg     string       `json:"msg,omitempty"`
	DataId  string       `json:"dataId,omitempty"`
	TaskId  string       `json:"taskId,omitempty"`
	Content string       `json:"content,omitempty"`
	Results []textresult `json:"results,omitempty"`
}
type textresult struct {
	Scene      string       `json:"scene,omitempty"`
	Suggestion string       `json:"suggestion,omitempty"`
	Label      string       `json:"label,omitempty"`
	Rate       float32      `json:"rate,omitempty"`
	Extras     interface{}  `json:"extras,omitempty"`
	Details    []textdetail `json:"details,omitempty"`
}
type textdetail struct {
	Label    string        `json:"label,omitempty"`
	Contexts []textcontext `json:"contexts,omitempty"`
}
type textcontext struct {
	Context  string `json:"context,omitempty"`
	LibName  string `json:"libName,omitempty"`
	RuleType string `json:"ruleType,omitempty"`
}

type ImgBizData struct {
	BizType string    `json:"bizType,omitempty"`
	Scenes  []string  `json:"scenes,omitempty"`
	Tasks   []ImgTask `json:"tasks,omitempty"`
}
type ImgTask struct {
	DataId string `json:"dataId,omitempty"`
	Url    string `json:"url,omitempty"`
}

type ImgResult struct {
	Code      int       `json:"code,omitempty"`
	Msg       string    `json:"msg,omitempty"`
	RequestId string    `json:"requestId,omitempty"`
	Data      []imgdata `json:"data,omitempty"`
}
type imgdata struct {
	Code    int               `json:"code,omitempty"`
	Msg     string            `json:"msg,omitempty"`
	DataId  string            `json:"dataId,omitempty"`
	TaskId  string            `json:"taskId,omitempty"`
	Url     string            `json:"url,omitempty"`
	Extras  map[string]string `json:"extras,omitempty"`
	Results []imgresult       `json:"results,omitempty"`
}
type imgresult struct {
	Scene      string         `json:"scene,omitempty"`
	Suggestion string         `json:"suggestion,omitempty"`
	Label      string         `json:"label,omitempty"`
	Rate       float32        `json:"rate,omitempty"`
	QrcodeData []string       `json:"qrcodeData,omitempty"`
	LogoData   []imglogData   `json:"logoData,omitempty"`
	SfaceData  []imgsfaceData `json:"sfaceData,omitempty"`
}
type imglogData struct {
	Type string  `json:"type,omitempty"`
	name string  `json:"name,omitempty"`
	X    float32 `json:"x,omitempty"`
	Y    float32 `json:"y,omitempty"`
	W    float32 `json:"w,omitempty"`
	H    float32 `json:"h,omitempty"`
}
type imgsfaceData struct {
	X     float32    `json:"x,omitempty"`
	Y     float32    `json:"y,omitempty"`
	W     float32    `json:"w,omitempty"`
	H     float32    `json:"h,omitempty"`
	Faces []imgfaces `json:"faces,omitempty"`
}
type imgfaces struct {
	Name string  `json:"name,omitempty"`
	Rate float32 `json:"rate,omitempty"`
	Id   string  `json:"id,omitempty"`
}

type VideoBizData struct {
	BizType string      `json:"bizType,omitempty"`
	Scenes  []string    `json:"scenes,omitempty"`
	Tasks   []VideoTask `json:"tasks,omitempty"`
}
type VideoTask struct {
	DataId      string       `json:"dataId,omitempty"`
	frames      []videoframe `json:"frames,omitempty"`
	framePrefix string       `json:"framePrefix,omitempty"`
	time        uint         `json:"time,omitempty"`
}
type videoframe struct {
	Url    string  `json:"url,omitempty"`
	Offset uint    `json:"offset,omitempty"`
	Rate   float32 `json:"rate,omitempty"`
}

type VideoResult struct {
	Code      int         `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	RequestId string      `json:"requestId,omitempty"`
	Data      []videodata `json:"data,omitempty"`
}
type videodata struct {
	Code      int           `json:"code,omitempty"`
	Msg       string        `json:"msg,omitempty"`
	DataId    string        `json:"dataId,omitempty"`
	RequestId string        `json:"requestId,omitempty"`
	TaskId    string        `json:"taskId,omitempty"`
	Results   []videoresult `json:"data,omitempty"`
}

type videoresult struct {
	Scene      string         `json:"scene,omitempty"`
	Label      string         `json:"label,omitempty"`
	Suggestion string         `json:"suggestion,omitempty"`
	Rate       float32        `json:"rate,omitempty"`
	Frames     []videoframe   `json:"frames,omitempty"`
	Extras     interface{}    `json:"extras,omitempty"`
	SfaceData  []imgsfaceData `json:"sfaceData,omitempty"`
}
