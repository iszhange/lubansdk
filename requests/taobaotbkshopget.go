package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-店铺搜索
// Link: https://open.taobao.com/api.htm?docId=24521&docType=2&scopeId=16516
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkShopGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkShopGet) New() *TaobaoTbkShopGet {
	r := new(TaobaoTbkShopGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkShopGet) Result(data []byte) (TaobaoTbkShopGetResponse, error) {
	var result TaobaoTbkShopGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkShopGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkShopGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkShopGet) SetFields(value string) {
	o.params["fields"] = value
}

func (o *TaobaoTbkShopGet) SetQ(value string) {
	o.params["q"] = value
}

func (o *TaobaoTbkShopGet) SetSort(value string) {
	o.params["sort"] = value
}

func (o *TaobaoTbkShopGet) SetIsTmall(value bool) {
	o.params["is_tmall"] = value
}

func (o *TaobaoTbkShopGet) SetStartCredit(value int) {
	o.params["start_credit"] = value
}

func (o *TaobaoTbkShopGet) SetEndCredit(value int) {
	o.params["end_credit"] = value
}

func (o *TaobaoTbkShopGet) SetStartCommissionRate(value int) {
	o.params["start_commission_rate"] = value
}

func (o *TaobaoTbkShopGet) SetEndCommissionRate(value int) {
	o.params["end_commission_rate"] = value
}

func (o *TaobaoTbkShopGet) SetStartTotalAction(value int) {
	o.params["start_total_action"] = value
}

func (o *TaobaoTbkShopGet) SetEndTotalAction(value int) {
	o.params["end_total_action"] = value
}

func (o *TaobaoTbkShopGet) SetStartAuctionCount(value int) {
	o.params["start_auction_count"] = value
}

func (o *TaobaoTbkShopGet) SetEndAuctionCount(value int) {
	o.params["end_auction_count"] = value
}

func (o *TaobaoTbkShopGet) SetPlatform(value int) {
	o.params["platform"] = value
}

func (o *TaobaoTbkShopGet) SetPageNo(value int) {
	o.params["page_no"] = value
}

func (o *TaobaoTbkShopGet) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkShopGet) GetMethod() string {
	return "taobao.tbk.shop.get"
}

func (o *TaobaoTbkShopGet) GetApiParams() Parameter {
	return o.params
}
