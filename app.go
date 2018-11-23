package main

import (
	"demo/router"
	"github.com/braintree/manners"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := router.InitRouter()
	//router.Run("0.0.0.0:8888")
	manners.ListenAndServe("0.0.0.0:8888", router)
}
