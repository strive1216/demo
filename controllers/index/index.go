package index

import (
	jwtauth "demo/middleware/JWT"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// func1: 处理最基本的GET
func Index(c *gin.Context) {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "index")
}
func Jwt(c *gin.Context) {
	j := &jwtauth.JWT{
		[]byte("test"),
	}
	claims := jwtauth.CustomClaims{
		1,
		"awh521",
		"1044176017@qq.com",
		jwt.StandardClaims{
			ExpiresAt: 15000, //time.Now().Add(24 * time.Hour).Unix()
			Issuer:    "test",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		c.Abort()
	}
	c.String(http.StatusOK, token+"---------------<br>")
	res, err := j.ParseToken(token)
	if err != nil {
		if err == jwtauth.TokenExpired {
			newToken, err := j.RefreshToken(token)
			if err != nil {
				c.String(http.StatusOK, err.Error())
			} else {
				c.String(http.StatusOK, newToken)
			}
		} else {
			c.String(http.StatusOK, err.Error())
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}
