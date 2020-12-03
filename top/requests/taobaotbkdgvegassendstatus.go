package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-所有订单查询
// Link: https://open.taobao.com/api.htm?docId=52958&docType=2&scopeId=21226
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgVegasSendStatus struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgVegasSendStatus) New() *TaobaoTbkDgVegasSendStatus {
	r := new(TaobaoTbkDgVegasSendStatus)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgVegasSendStatus) Result(data []byte) (TaobaoTbkDgVegasSendStatusResponse, error) {
	var result TaobaoTbkDgVegasSendStatusResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgVegasSendStatusResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgVegasSendStatusResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgVegasSendStatus) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkDgVegasSendStatus) GetMethod() string {
	return "taobao.tbk.dg.vegas.send.status"
}

func (o *TaobaoTbkDgVegasSendStatus) GetApiParams() Parameter {
	return o.params
}
