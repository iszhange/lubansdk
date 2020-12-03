package requests

import (
	"encoding/json"
)

// 刷新Access Token
// Link: https://open.taobao.com/API.htm?docId=25387&docType=2
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTopAuthTokenRefresh struct {
	params Parameter // 参数
}

func (o TaobaoTopAuthTokenRefresh) New() *TaobaoTopAuthTokenRefresh {
	r := new(TaobaoTopAuthTokenRefresh)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTopAuthTokenRefresh) Result(data []byte) (TaobaoTopAuthTokenRefreshResponse, error) {
	var result TaobaoTopAuthResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTopAuthTokenRefreshResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTopAuthTokenRefreshResponse{}, API_RESPONSE_ERROR
	}

	var token TaobaoTopAuthTokenResult
	err = json.Unmarshal([]byte(result.TokenResult), &token)
	if err != nil {
		return TaobaoTopAuthTokenRefreshResponse{}, err
	}

	var response TaobaoTopAuthTokenRefreshResponse
	response.RequestID = result.RequestID
	response.TokenResult = token

	return response, nil
}

func (o *TaobaoTopAuthTokenRefresh) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTopAuthTokenRefresh) GetMethod() string {
	return "taobao.top.auth.token.refresh"
}

func (o *TaobaoTopAuthTokenRefresh) GetApiParams() Parameter {
	return o.params
}
