package aliyun

import (
	"demo/lib"
	"demo/lib/aliyun"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type param struct {
	Img   string   `json:"img,omitempty"`
	Imgs  []string `json:"imgs,omitempty"`
	Text  string   `json:"text,omitempty"`
	Texts []string `json:"texts,omitempty"`
}

var textscene = []string{"antispam"}
var imgscene = []string{"porn", "terrorism"}

func InspectText(c *gin.Context) {
	var info param
	err := c.BindJSON(&info)
	if err != nil {
		fmt.Println(err.Error())
	}
	var data []byte
	if info.Text != "" {
		slice1 := make([]string, 0)
		n := 4000
		ll := info.Text
		for n < len([]rune(ll)) {
			rs := []rune(ll)
			m := lib.Substr2(ll, 0, n)
			slice1 = append(slice1, m)
			ll = lib.Substr2(string(rs), n, len(rs))
		}
		slice1 = append(slice1, ll)
		data = aliyun.InspectText(slice1, "Green", textscene)
	} else {
		data = aliyun.InspectText(info.Texts, "Green", textscene)
	}

	fmt.Println(string(data))
	m := aliyun.TextResult{}
	json.Unmarshal(data, &m)
	var aa bool = false
	if m.Code == 200 {
		aa = true
	}
	c.SecureJSON(http.StatusOK, gin.H{
		"success": aa,
		"data":    m.Data,
		"code":    m.Code,
		"msg":     m.Msg,
	})
}

func InspectImg(c *gin.Context) {
	var info param
	c.BindJSON(&info)
	var data []byte
	if info.Img != "" {
		data = aliyun.InspectImg([]string{info.Img}, "Green", imgscene)
	} else {
		data = aliyun.InspectImg(info.Imgs, "Green", imgscene)
	}
	m := aliyun.ImgResult{}
	json.Unmarshal(data, &m)
	var aa bool = false
	if m.Code == 200 {
		aa = true
	}
	c.SecureJSON(http.StatusOK, gin.H{
		"success": aa,
		"data":    m.Data,
		"code":    m.Code,
		"msg":     m.Msg,
	})
}

func Inspect(c *gin.Context) {
	var info param
	c.BindJSON(&info)

	var wg sync.WaitGroup
	wg.Add(2)
	var (
		aa []byte
		bb []byte
	)
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("文字", time.Now().UnixNano())
		aa = aliyun.InspectText([]string{info.Text}, "Green", textscene)
		wg.Done()
	}()
	go func() {
		fmt.Println("图片", time.Now().UnixNano())
		bb = aliyun.InspectImg([]string{info.Img}, "Green", imgscene)
		wg.Done()
	}()
	wg.Wait()
	n := aliyun.TextResult{}
	json.Unmarshal(aa, &n)
	m := aliyun.ImgResult{}
	json.Unmarshal(bb, &m)

	c.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"text":    n,
		"img":     m,
	})

}
