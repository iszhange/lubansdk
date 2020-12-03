package requests

import (
	"encoding/json"
)

// 淘宝客-公用-店铺关联推荐
// Link: https://open.taobao.com/api.htm?docId=24522&docType=2&scopeId=16292
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkShopRecommendGet struct {
	params Parameter
}

func (o TaobaoTbkShopRecommendGet) New() *TaobaoTbkShopRecommendGet {
	r := new(TaobaoTbkShopRecommendGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkShopRecommendGet) Result(data []byte) (TaobaoTbkShopRecommendGetResponse, error) {
	var result TaobaoTbkShopRecommendGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkShopRecommendGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkShopRecommendGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkShopRecommendGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkShopRecommendGet) GetMethod() string {
	return "taobao.tbk.shop.recommend.get"
}

func (o *TaobaoTbkShopRecommendGet) GetApiParams() Parameter {
	return o.params
}
