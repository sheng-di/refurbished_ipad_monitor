package main

import (
	"applebuyutil"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"wechat"
)

// 60 秒内只发送一次
const ININTEVAL = 60

var lastTime int64
var canSend bool

func GetInfo() (*applebuyutil.Info, error) {
	apiUrl := "https://www.apple.com.cn/shop/refurbished/ipad"
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	// 解析结果
	var result applebuyutil.Info
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	arr := strings.Split(string(body), "window.REFURB_GRID_BOOTSTRAP = ")
	if len(arr) != 2 {
		return nil, errors.New("解析失败 1")
	}
	arr = strings.Split(arr[1], ";")
	if len(arr) < 2 {
		return nil, errors.New("解析失败 2")
	}
	err = json.Unmarshal([]byte(arr[0]), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func sendItem(item applebuyutil.Tile, index int, money int) {
	content := fmt.Sprintf("%d, %s, 【%v】, 价格低于 %d", index, item.Title, item.Price.SeoPrice, money)
	log.Println(content)
	// 推送
	if canSend {
		_, err := wechat.SendText(content)
		lastTime = time.Now().Unix()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("推送完毕")
	}
}

func Monitor() {
	/**
	监视：
	1、Air, 小于4000元
	2、Pro, 小于5000元
	*/
	//var sends []string
	result, err := GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	// 价格范围
	var content string
	airMoney := 4000
	proMoney := 5000
	log.Printf("一共有%d项翻新产品", len(result.Tiles))
	canSend = time.Now().Unix()-lastTime >= ININTEVAL
	for index, item := range result.Tiles {
		if strings.Contains(item.Title, "iPad Air") && item.Price.SeoPrice < float64(airMoney) {
			sendItem(item, index, airMoney)
		} else if strings.Contains(item.Title, "iPad Pro") && item.Price.SeoPrice < float64(proMoney) {
			sendItem(item, index, proMoney)
		} else {
			// 其他的，默认打印在屏幕上即可
			content = fmt.Sprintf("%d, %s, 【%v】", index, item.Title, item.Price.SeoPrice)
			log.Println(content)
		}
	}
}

func main() {
	for {
		Monitor()
		time.Sleep(time.Second * 5)
	}
}
