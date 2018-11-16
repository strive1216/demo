package account

import (
	"demo/models/account"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"time"
)

type user = account.Account

var (
	Account =  account.Account{}
)

// func1: 处理最基本的GET
func Func1(c   *gin.Context) {
	name := c.DefaultQuery("name", "中国")
	Id := c.Query("Id")
	
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "hello %v %s", name, Id)
}

// func2: 处理最基本的POST
func Func2(c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	name := c.PostForm("name")
	password := c.DefaultPostForm("password", "123456")

	var userinfo user
	err := c.BindJSON(&userinfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "404"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": name, "password": password, "userinfo": userinfo})

	//struct 转 json 串
	jsons, errs := json.Marshal(userinfo) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(jsons))
}

func Func3(c *gin.Context) {
	var person user
	if err := c.ShouldBindJSON(&person); err == nil {
		c.JSON(http.StatusOK, person)
	} else {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 404})
	}
}

func Func4(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	//c.String(http.StatusOK, string(buf[0:n]))
	c.JSONP(http.StatusOK, string(buf[0:n]))
	//resp := map[string]string{"hello": "world"}
	//c.JSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	var userInfo user
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "404"})
		return
	}
	exist := Account.IsExist(bson.M{"username": userInfo.Username})
	if exist {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": userInfo})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 403, "data": userInfo})
	}
}

func ListAccount(c *gin.Context) {
	exist, err := Account.FindAllAccount()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": exist})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 403, "data": ""})
	}
}

func Register(c *gin.Context) {
	var userInfo user
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "404"})
		return
	}
	userInfo.Id = bson.NewObjectId()
	userInfo.CreatedAt = time.Now()
	userInfo.ModifiedAt = time.Now()
	err_ := Account.Insert(userInfo)
	if err_ == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": userInfo})
	} else {
		fmt.Println(err_.Error())
		c.JSON(http.StatusOK, gin.H{"code": 403, "data": userInfo})
	}

}

func Info(c *gin.Context) {
	id := c.Param("id")
	exist, err := Account.FindById(id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": exist})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 404, "data": ""})
	}
}

func Update(c *gin.Context) {
	var userInfo user
	err := c.BindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "404"})
		return
	}
	userInfo.Id = bson.ObjectIdHex(c.Param("id"))
	userInfo.ModifiedAt = time.Now()
	_err := Account.UpdateAccount(userInfo)
	if _err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200})
	} else {
		fmt.Println(_err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 404})
	}
}

func Remove(c *gin.Context) {
	id := c.Param("id")
	err := Account.RemoveAccount(id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 404})
	}
}
