package requests

import (
	"encoding/json"
)

// 淘宝客-公用-长链转短链
// Link: https://open.taobao.com/api.htm?docId=27832&docType=2&scopeId=12340
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkSpreadGet struct {
	params []Parameter
}

func (o TaobaoTbkSpreadGet) New() *TaobaoTbkSpreadGet {
	r := new(TaobaoTbkSpreadGet)
	r.params = make([]Parameter, 0)
	return r
}

func (o TaobaoTbkSpreadGet) Result(data []byte) (TaobaoTbkSpreadGetResponse, error) {
	var result TaobaoTbkSpreadGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkSpreadGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkSpreadGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkSpreadGet) SetParam(key, value string) {
	o.params = append(o.params, Parameter{key: value})
}

func (o *TaobaoTbkSpreadGet) GetMethod() string {
	return "taobao.tbk.spread.get"
}

func (o *TaobaoTbkSpreadGet) GetApiParams() Parameter {
	p, _ := json.Marshal(o.params)
	var result = make(Parameter)
	result["requests"] = string(p)

	return result
}
