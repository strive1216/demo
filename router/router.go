package router

import (
	"demo/controllers/account"
	"demo/controllers/index"
	"demo/controllers/note"
	"demo/lib/aliyun"
	"demo/middleware"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/resources", "./resources")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.GET("/", index.Index)
	//
	router.GET("/jwt", index.Jwt)
	router.GET("/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	r1 := router.Group("/account").Use(middleware.SetHeaderJSON())
	{
		r1.GET("/", account.Func1)
		//r1.POST("/",account.Func2)
		//r1.POST("/a",account.Func3)
		//r1.POST("/b",account.Func4)
		r1.POST("/login", account.Login)
		r1.GET("/list", account.ListAccount)
		r1.POST("/reg", account.Register)
		r1.GET("/info/:id", account.Info)
		r1.DELETE("/rem/:id", account.Remove)
		r1.PUT("/update/:id", account.Update)
		r1.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
	}

	r2 := router.Group("/note").Use(middleware.SetHeaderJSON())
	{
		r2.POST("/", note.Insert)
		r2.GET("/:id", note.Get)
		r2.GET("/", note.GetAll)
		//r2.POST("/login", account.Login)
		//r2.GET("/info/:id", account.Info)
	}
	r3 := router.Group("/aliyun").Use(middleware.SetHeaderJSON())
	{
		r3.POST("/text", func(c *gin.Context) {
			//buf := make([]byte, 1024)
			//n, _ := c.Request.Body.Read(buf)
			//fmt.Println(string(buf[0:n]))

			text := c.PostForm("aa")
			data := aliyun.InspectText(text, 4000)
			fmt.Println(string(data))
			m := aliyun.TextResult{}
			json.Unmarshal(data, &m)
			c.SecureJSON(http.StatusOK, gin.H{
				"success": true,
				"data":    m.Data,
				"code":    m.Code,
				"msg":     m.Msg,
			})
		})
		r3.POST("/img", func(c *gin.Context) {
			img := c.Request.PostForm.Get("img")
			c.String(http.StatusOK, string(aliyun.InspectImg([]string{img})))
		})
	}
	return router
}
