package main

import (
	"fmt"
	"github.com/CBDlkl/gin"
	"time"
	"./logic"
)

var ctr = new(logic.Ctrip)

func init() {
	ctr.Login()
	go ctr.Heartbeat()
	fmt.Println("init success ...")
}

func main() {
	r := setupRoter()
	r.Run(":5005")
}

func setupRoter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", func(context *gin.Context) {
		for {
			if ctr.Authorization == "" {
				ctr.Login()
			}
			time.Sleep(2 * time.Second)
		}
		context.String(200, ctr.Authorization)
	})

	return router
}
