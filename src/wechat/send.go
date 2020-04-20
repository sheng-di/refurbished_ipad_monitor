package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const SERVER_CHAN_KEY = "从Server酱获取: "

type SendResult struct {
	ErrNo   int    `json:"errno"`
	ErrMsg  string `json:"errmsg"`
	Dataset string `json:"dataset"`
}

func Send(text string, description string) (*SendResult, error) {
	apiUrl := fmt.Sprintf("https://sc.ftqq.com/%s.send?text=%s&desp=%s", SERVER_CHAN_KEY, url.QueryEscape(text), url.QueryEscape(description))
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("结果为空")
	} else {
		defer res.Body.Close()
	}
	// 解析结果
	var result SendResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

func SendWithTimestamp(text string, description string) (*SendResult, error) {
	description = fmt.Sprintf("%s\n%v", description, time.Now().Unix())
	return Send(text, description)
}

func SendMessage(text string, description string) (*SendResult, error) {
	text = fmt.Sprintf("mpGo，%s", text)
	return SendWithTimestamp(text, description)
}
func SendText(text string) (*SendResult, error) {
	return SendWithTimestamp(text, "")
}
