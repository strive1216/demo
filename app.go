package main

import (
	"demo/lib/aliyun"
	"demo/router"
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func main() {
	fmt.Println("开始", time.Now().UnixNano())
	go ialiyun()
	fmt.Println("mian结束", time.Now().UnixNano())

	gin.SetMode(gin.ReleaseMode)
	router := router.InitRouter()
	//router.Run("0.0.0.0:8888")
	manners.ListenAndServe("0.0.0.0:8888", router)

}

func ialiyun() {
	var wg sync.WaitGroup
	wg.Add(2)
	var (
		aa []byte
		bb []byte
	)
	go func() {
		fmt.Println("文字")
		time.Sleep(time.Duration(2) * time.Second)
		aa = aliyun.InspectText("123456789abcdefjhijklnmopqrstuvwsyz 做爱", 33)
		fmt.Println("文字", time.Now().UnixNano())
		wg.Done()
	}()

	go func() {
		fmt.Println("图片")
		bb = aliyun.InspectImg([]string{"https://cdn.stonedrums.com.cn/201810111539238297413.jpg"})
		fmt.Println("图片", time.Now().UnixNano())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(string(aa))
	//m :=aliyun.Textget{}
	//json.Unmarshal(aa, &m)
	//fmt.Println(m.Code)

	fmt.Println(string(bb))
	fmt.Println("结束", time.Now().UnixNano())
}
