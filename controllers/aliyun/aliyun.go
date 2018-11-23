package aliyun

import (
	"demo/lib/aliyun"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type param struct {
	img   string   `json:"img,omitempty"`
	imgs  []string `json:"imgs,omitempty"`
	text  string   `json:"text,omitempty"`
	texts []string `json:"texts,omitempty"`
}

func InspectText(c *gin.Context) {
	var info param
	c.BindJSON(&info)
	data := aliyun.InspectText(info.text, 4000)
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
	if info.img != "" {
		data = aliyun.InspectImg([]string{info.img})
	} else {
		data = aliyun.InspectImg(info.imgs)
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
	text := c.PostForm("aa")
	img := c.PostForm("img")

	var wg sync.WaitGroup
	wg.Add(2)
	var (
		aa []byte
		bb []byte
	)
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("文字", time.Now().UnixNano())
		aa = aliyun.InspectText(text, 33)
		wg.Done()
	}()
	go func() {
		fmt.Println("图片", time.Now().UnixNano())
		bb = aliyun.InspectImg([]string{img})
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
