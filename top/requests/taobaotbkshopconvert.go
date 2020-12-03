package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-店铺链接转换
// Link: https://open.taobao.com/api.htm?docId=24523&docType=2&scopeId=11653
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkShopConvert struct {
	params Parameter // 参数
}

func (o TaobaoTbkShopConvert) New() *TaobaoTbkShopConvert {
	r := new(TaobaoTbkShopConvert)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkShopConvert) Result(data []byte) (TaobaoTbkShopConvertResponse, error) {
	var result TaobaoTbkShopConvertResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkShopConvertResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkShopConvertResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkShopConvert) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkShopConvert) GetMethod() string {
	return "taobao.tbk.shop.convert"
}

func (o *TaobaoTbkShopConvert) GetApiParams() Parameter {
	return o.params
}
