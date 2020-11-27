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

func (o *TaobaoTbkOrderDetailsGet) SetQueryType(value int) {
	o.params["query_type"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetOrderScene(value int) {
	o.params["order_scene"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetPositionIndex(value string) {
	o.params["position_index"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetMemberType(value int) {
	o.params["member_type"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetTkStatus(value int) {
	o.params["tk_status"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetStartTime(value string) {
	o.params["start_time"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetEndTime(value string) {
	o.params["end_time"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetJumpType(value int) {
	o.params["jump_type"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetPageNo(value int) {
	o.params["page_no"] = value
}

func (o *TaobaoTbkOrderDetailsGet) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkOrderDetailsGet) GetMethod() string {
	return "taobao.tbk.order.details.get"
}

func (o *TaobaoTbkOrderDetailsGet) GetApiParams() Parameter {
	return o.params
}
