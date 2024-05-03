package http

import (
	"crypto/tls"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client
type Client struct {
	url    string
	client *http.Client
}

// NewClient 初始化客户端
func NewClient(url string) *Client {
	return &Client{
		url: url,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10, // 每台主机保持的最大空闲连接
				MaxConnsPerHost:     10, // 限制每个主机的连接总数
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

// httpRequest http请求
func (c *Client) HttpRequest(methed, params string) (resBody []byte, err error) {
	// 初始化请求
	body := strings.NewReader(params)
	req, err := http.NewRequest("POST", c.url+methed, body)
	if err != nil {
		return nil, errors.Wrap(err, "Http NewRequest")
	}
	// 执行请求
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TRON-PRO-API-KEY","c1c9a2da-8fd0-40d7-9599-d55129f43c2d")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Client Do")
	}
	defer res.Body.Close()
	// 接收返回结果
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}
	return resBody, nil
}

// httpRequest http请求
func (c *Client) HttpGetRequest(methed, params string) (resBody []byte, err error) {

	// 初始化请求
	req, err := http.NewRequest("GET", c.url+methed+params, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Http NewRequest")
	}

	// 执行请求
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TRON-PRO-API-KEY","c1c9a2da-8fd0-40d7-9599-d55129f43c2d")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Client Do")
	}
	defer res.Body.Close()

	// 接收返回结果
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}
	return resBody, nil
}
