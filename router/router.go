package router

import (
	"demo/controllers/account"
	"demo/controllers/aliyun"
	"demo/controllers/index"
	"demo/controllers/note"
	"demo/middleware"
	"demo/middleware/JWT"
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
		r1.GET("/info/:id", JWT.JWTAuth(), account.Info)
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
		r3.POST("/text", aliyun.InspectText)
		r3.POST("/img", aliyun.InspectImg)
		r3.POST("/all", aliyun.Inspect)
	}
	return router
}
