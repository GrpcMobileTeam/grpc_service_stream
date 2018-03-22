package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"time"
	"encoding/json"
	"github.com/CBDlkl/gin"
)

var (
	Authorization string
)

func init() {
	login()
	go heartbeat()
	fmt.Println("init success ...")
}

func main() {
	r := setupRoter()
	r.Run(":5005")
}

func setupRoter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", func(context *gin.Context) {
		context.String(200, Authorization)
	})

	return router
}

// 登录
func login() {
	for i := 1; i < 4; i++ {
		resp, err := http.Post("http://172.16.5.68:8080/v1/token", "application/x-www-form-urlencoded",
			strings.NewReader(fmt.Sprintf("u=%s&p=%s", "ceshi1", "xc123456")))
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		var jsonData map[string]interface{}
		json.Unmarshal(body, &jsonData)

		if jsonData["status"] == "ok" {
			// set auto
			fmt.Println("获取到Authorization => ", jsonData["token"].(string))
			Authorization = jsonData["token"].(string)
			break
		} else {
			fmt.Errorf("第%d次请求登录", i)
		}
		time.Sleep(3 * time.Second)
	}
}

func heartbeat() {
	for {
		defer func() {
			if err := recover(); err != nil {
				fmt.Errorf("header beat err, err info => %#v", err)
			}
		}()

		time.Sleep(5 * time.Second)

		client := &http.Client{}

		req, err := http.NewRequest("POST", "https://can.ctrip.com/api/user/getfavoritelist", strings.NewReader("{\"Action\":\"GETLIST\",\"CustomCurrency\":\"CNY\"}"))
		if err != nil {
			fmt.Errorf("post err, err info => %#v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Host", "can.ctrip.com")
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("Origin", "https://can.ctrip.com")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
		req.Header.Set("Authorization", Authorization)

		resp, _ := client.Do(req)
		defer resp.Body.Close()

		if resp.StatusCode == 401 {
			login()
		}
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(resp.StatusCode, string(body))
	}
}

//func httpPost(url string, data interface{}) (string, error) {
//	client := &http.Client{}
//
//	req, err := http.NewRequest("POST", url, strings.NewReader(data.(string)))
//	if err != nil {
//		fmt.Errorf("post err, err info => %#v", err)
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//
//	resp, err := client.Do(req)
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Errorf("post read response err , err info => %#v", err)
//	}
//
//	return string(body), err
//}
