package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-权益物料精选
// Link: https://open.taobao.com/api.htm?docId=52700&docType=2&scopeId=16518
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgOptimusPromotion struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgOptimusPromotion) New() *TaobaoTbkDgOptimusPromotion {
	r := new(TaobaoTbkDgOptimusPromotion)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgOptimusPromotion) Result(data []byte) (TaobaoTbkDgOptimusPromotionResponse, error) {
	var result TaobaoTbkDgOptimusPromotionResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgOptimusPromotionResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgOptimusPromotionResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgOptimusPromotion) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkDgOptimusPromotion) SetPageNum(value int) {
	o.params["page_num"] = value
}

func (o *TaobaoTbkDgOptimusPromotion) SetAdzoneId(value string) {
	o.params["adzone_id"] = value
}

func (o *TaobaoTbkDgOptimusPromotion) SetPromotionId(value int) {
	o.params["promotion_id"] = value
}

func (o *TaobaoTbkDgOptimusPromotion) GetMethod() string {
	return "taobao.tbk.dg.optimus.promotion"
}

func (o *TaobaoTbkDgOptimusPromotion) GetApiParams() Parameter {
	return o.params
}
