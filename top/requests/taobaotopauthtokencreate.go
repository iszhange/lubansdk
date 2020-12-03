package requests

import (
	"encoding/json"
)

// 获取Access Token
// Link: https://open.taobao.com/API.htm?docId=25388&docType=2
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTopAuthTokenCreate struct {
	params Parameter // 参数
}

func (o TaobaoTopAuthTokenCreate) New() *TaobaoTopAuthTokenCreate {
	r := new(TaobaoTopAuthTokenCreate)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTopAuthTokenCreate) Result(data []byte) (TaobaoTopAuthTokenCreateResponse, error) {
	var result TaobaoTopAuthResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTopAuthTokenCreateResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTopAuthTokenCreateResponse{}, API_RESPONSE_ERROR
	}

	var token TaobaoTopAuthTokenResult
	err = json.Unmarshal([]byte(result.TokenResult), &token)
	if err != nil {
		return TaobaoTopAuthTokenCreateResponse{}, err
	}

	var response TaobaoTopAuthTokenCreateResponse
	response.RequestID = result.RequestID
	response.TokenResult = token

	return response, nil
}

func (o *TaobaoTopAuthTokenCreate) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTopAuthTokenCreate) GetMethod() string {
	return "taobao.top.auth.token.create"
}

func (o *TaobaoTopAuthTokenCreate) GetApiParams() Parameter {
	return o.params
}
