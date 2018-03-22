package main

import (
	"fmt"
	"github.com/CBDlkl/gin"
)

var ctrip = new(Ctrip)

func init() {
	ctrip.login()
	go ctrip.heartbeat()
	fmt.Println("init success ...")
}

func main() {
	r := setupRoter()
	r.Run(":5005")
}

func setupRoter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", func(context *gin.Context) {
		context.String(200, ctrip.Authorization)
	})

	return router
}
