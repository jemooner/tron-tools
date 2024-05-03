package jsonrpc

import (
	"encoding/json"
	"fmt"
	"sync"
	"tron-tools/rpcclient/http"
)

type Http struct {
	rpc *http.Client
}

var (
	httpRPC  *Http
	httpOnce sync.Once
)

// NewETHTP
func NewETHTP(host string) *Http {
	httpOnce.Do(func() {
		httpRPC = &Http{http.NewClient(host)}
	})
	return httpRPC
}

// 根据区块号获取区块信息struct
type Num struct {
	Num uint64 `json:"num"`
}

// 根据区块号获取区块信息struct
type HashId struct {
	Value string `json:"value"`
}

// GetNowBlock 获取最新区块
func (tron *Http) GetNowBlock() ([]byte, error) {
	params := NewHttpParams("")
	resBody, err := tron.rpc.HttpRequest("/wallet/getnowblock", params)
	if err != nil {
		return nil, err
	}
	if resBody == nil {
		return nil, fmt.Errorf("resBody is null")
	}
	return resBody, err
}

// GetBlockByNumber 根据块号获取区块信息
func (tron *Http) GetBlockByNumber(number uint64) ([]byte, error) {
	var num Num
	num.Num = number
	params := NewHttpParams(num)
	resBody, err := tron.rpc.HttpRequest("/wallet/getblockbynum", params)
	if err != nil {
		return nil, err
	}
	if resBody == nil {
		return nil, fmt.Errorf("resBody is null")
	}
	return resBody, nil
}

func NewHttpParams(args interface{}) string {
	rb, _ := json.Marshal(args)
	return string(rb)
}
