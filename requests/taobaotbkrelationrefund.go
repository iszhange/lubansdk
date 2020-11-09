package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-所有订单查询
// Link: https://open.taobao.com/api.htm?docId=40121&docType=2&scopeId=16175
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkRelationRefund struct {
	params Parameter // 参数
}

func (o TaobaoTbkRelationRefund) New() *TaobaoTbkRelationRefund {
	r := new(TaobaoTbkRelationRefund)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkRelationRefund) Result(data []byte) (TaobaoTbkRelationRefundResponse, error) {
	var result TaobaoTbkRelationRefundResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkRelationRefundResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkRelationRefundResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkRelationRefund) SetBizType(value int) {
	o.params["biz_type"] = value
}

func (o *TaobaoTbkRelationRefund) SetStartTime(value string) {
	o.params["start_time"] = value
}

func (o *TaobaoTbkRelationRefund) SetRefundType(value int) {
	o.params["refund_type"] = value
}

func (o *TaobaoTbkRelationRefund) SetSearchType(value int) {
	o.params["search_type"] = value
}

func (o *TaobaoTbkRelationRefund) SetPageNo(value int) {
	o.params["page_no"] = value
}

func (o *TaobaoTbkRelationRefund) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkRelationRefund) GetMethod() string {
	return "taobao.tbk.relation.refund"
}

func (o *TaobaoTbkRelationRefund) GetApiParams() Parameter {
	p, _ := json.Marshal(o.params)
	var result = make(Parameter)
	result["search_option"] = string(p)

	return result
}
