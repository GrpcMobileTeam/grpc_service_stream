package main

import (
	"fmt"
	"time"
)

var PostQueue = make(chan string, 5)

func main() {
	go goQueue()
	write()

	select {}
}

func write() {
	fmt.Println("begin...")
	time.Sleep(5)
	PostQueue <- "ok"
}

func goQueue() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error to chan put.")
		}
	}()

	fmt.Println("goroutine running...")
	select {
	case p := <-PostQueue:
		fmt.Println("协程捕获", p)
	}
	fmt.Println("goroutine end")
}
