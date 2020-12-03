package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-淘礼金创建
// Link: https://open.taobao.com/api.htm?docId=40173&docType=2&scopeId=15029
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgVegasTljCreate struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgVegasTljCreate) New() *TaobaoTbkDgVegasTljCreate {
	r := new(TaobaoTbkDgVegasTljCreate)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgVegasTljCreate) Result(data []byte) (TaobaoTbkDgVegasTljCreateResponse, error) {
	var result TaobaoTbkDgVegasTljCreateResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgVegasTljCreateResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgVegasTljCreateResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgVegasTljCreate) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkDgVegasTljCreate) GetMethod() string {
	return "taobao.tbk.dg.vegas.tlj.create"
}

func (o *TaobaoTbkDgVegasTljCreate) GetApiParams() Parameter {
	return o.params
}
