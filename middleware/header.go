package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetHeaderJSON() gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("ssss")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8;")
	}
}
