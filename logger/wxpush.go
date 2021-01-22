package logger

import (
	"fmt"
	"net/http"
	"net/url"
)

type WxPush struct {
	host  string // 微信报警地址
	email string // 接收报请者email
}

func NewWxPush(host string, email string) *WxPush {
	return &WxPush{
		host:  host,
		email: email,
	}
}

func (this *WxPush) Write(p []byte) (n int, err error) {
	msg := string(p)
	if len(msg) == 0 || len(this.host) == 0 || len(this.email) == 0 {
		return len(msg), nil
	}
	go this.push(msg)
	return len(msg), nil
}

// Sync 实现 zap 接口
func (this *WxPush) Sync() error {
	return nil
}

// 发送推送
func (this *WxPush) push(msg string) error {
	monitorUrl := fmt.Sprintf("%s?email=%s&msg=%s", this.host, this.email, url.QueryEscape(msg))
	resp, _ := http.Get(monitorUrl)
	//defer resp.Body.Close()
	fmt.Println(resp)
	return nil
}
