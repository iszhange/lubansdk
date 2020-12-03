package requests

import (
	"encoding/json"
)

// 淘宝客-公用-阿里妈妈推广券详情查询
// Link: https://open.taobao.com/api.htm?docId=31106&docType=2&scopeId=16189
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkCouponGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkCouponGet) New() *TaobaoTbkCouponGet {
	r := new(TaobaoTbkCouponGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkCouponGet) Result(data []byte) (TaobaoTbkCouponGetResponse, error) {
	var result TaobaoTbkCouponGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkCouponGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkCouponGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkCouponGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkCouponGet) GetMethod() string {
	return "taobao.tbk.coupon.get"
}

func (o *TaobaoTbkCouponGet) GetApiParams() Parameter {
	return o.params
}
