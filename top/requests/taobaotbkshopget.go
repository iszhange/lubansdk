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

func (o *TaobaoTbkShopGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkShopGet) GetMethod() string {
	return "taobao.tbk.shop.get"
}

func (o *TaobaoTbkShopGet) GetApiParams() Parameter {
	return o.params
}
