package requests

import (
	"encoding/json"
)

// 淘宝客-公用-淘宝客商品详情查询(简版)
// Link: https://open.taobao.com/API.htm?docId=24518&docType=2
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkItemInfoGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkItemInfoGet) New() *TaobaoTbkItemInfoGet {
	r := new(TaobaoTbkItemInfoGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkItemInfoGet) Result(data []byte) (TaobaoTbkItemInfoGetResponse, error) {
	var result TaobaoTbkItemInfoGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkItemInfoGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkItemInfoGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkItemInfoGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkItemInfoGet) GetMethod() string {
	return "taobao.tbk.item.info.get"
}

func (o *TaobaoTbkItemInfoGet) GetApiParams() Parameter {
	return o.params
}
