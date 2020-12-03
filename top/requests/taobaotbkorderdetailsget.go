package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-所有订单查询
// Link: https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkOrderDetailsGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkOrderDetailsGet) New() *TaobaoTbkOrderDetailsGet {
	r := new(TaobaoTbkOrderDetailsGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkOrderDetailsGet) Result(data []byte) (TaobaoTbkOrderDetailsGetResponse, error) {
	var result TaobaoTbkOrderDetailsGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkOrderDetailsGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkOrderDetailsGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkOrderDetailsGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkOrderDetailsGet) GetMethod() string {
	return "taobao.tbk.order.details.get"
}

func (o *TaobaoTbkOrderDetailsGet) GetApiParams() Parameter {
	return o.params
}
