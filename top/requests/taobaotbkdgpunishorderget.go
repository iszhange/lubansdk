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

func (o *TaobaoTbkDgPunishOrderGet) SetSiteId(value string) {
	o.params["site_id"] = value
}

func (o *TaobaoTbkDgPunishOrderGet) SetSpan(value string) {
	o.params["span"] = value
}

func (o *TaobaoTbkDgPunishOrderGet) SetRelationId(value string) {
	o.params["relation_id"] = value
}

func (o *TaobaoTbkDgPunishOrderGet) SetTbTradeId(value string) {
	o.params["tb_trade_id"] = value
}

func (o *TaobaoTbkDgPunishOrderGet) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkDgPunishOrderGet) SetPageNo(value int) {
	o.params["page_no"] = value
}

func (o *TaobaoTbkDgPunishOrderGet) SetStartTime(value string) {
	o.params["start_time"] = value
}

func (o *TaobaoTbkDgPunishOrderGet) SetAdzoneId(value string) {
	o.params["adzone_id"] = value
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
