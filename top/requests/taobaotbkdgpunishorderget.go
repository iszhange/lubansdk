package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-处罚订单查询
// Link: https://open.taobao.com/api.htm?docId=42050&docType=2&scopeId=15736
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgPunishOrderGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgPunishOrderGet) New() *TaobaoTbkDgPunishOrderGet {
	r := new(TaobaoTbkDgPunishOrderGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgPunishOrderGet) Result(data []byte) (TaobaoTbkDgPunishOrderGetResponse, error) {
	var result TaobaoTbkDgPunishOrderGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgPunishOrderGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgPunishOrderGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgPunishOrderGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkDgPunishOrderGet) GetMethod() string {
	return "taobao.tbk.dg.punish.order.get"
}

func (o *TaobaoTbkDgPunishOrderGet) GetApiParams() Parameter {
	p, _ := json.Marshal(o.params)
	var result = make(Parameter)
	result["af_order_option"] = string(p)

	return result
}
