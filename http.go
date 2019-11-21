package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewClient(connTimeout, timeout int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(connTimeout)) // 设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(time.Second * time.Duration(timeout))) // 设置发送接收数据超时
				return c, nil
			},
		},
	}
}

// 发送post表单请求
func HttpPostForm(httpUrl string, data map[string]interface{}, connTimeout, timeout int) (*http.Response, error) {
	c := NewClient(connTimeout, timeout)
	params := url.Values{}
	for k, v := range data {
		params[k] = []string{fmt.Sprintf("%v", v)}
	}
	return c.PostForm(httpUrl, params)
}

// 发送GET请求
func HttpGet(httpUrl string, data map[string]interface{}, connTimeout, timeout int) (*http.Response, error) {
	c := NewClient(connTimeout, timeout)
	params := []string{}
	for k, v := range data {
		params = append(params, fmt.Sprintf("%v=%v", k, v))
	}
	if len(params) > 0 {
		httpUrl = fmt.Sprintf("%v?%v", httpUrl, strings.Join(params, "&"))
	}

	return c.Get(httpUrl)
}

// post json
func HttpPostJson(httpUrl string, data map[string]interface{}, connTimeout, timeout int) (*http.Response, error) {
	c := NewClient(connTimeout, timeout)
	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", httpUrl, bytes.NewReader(bytesData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return c.Do(req)
}
