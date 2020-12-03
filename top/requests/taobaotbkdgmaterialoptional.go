package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-物料搜索
// Link: https://open.taobao.com/api.htm?docId=35896&docType=2&scopeId=16516
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgMaterialOptional struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgMaterialOptional) New() *TaobaoTbkDgMaterialOptional {
	r := new(TaobaoTbkDgMaterialOptional)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgMaterialOptional) Result(data []byte) (TaobaoTbkDgMaterialOptionalResponse, error) {
	var result TaobaoTbkDgMaterialOptionalResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgMaterialOptionalResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgMaterialOptionalResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgMaterialOptional) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkDgMaterialOptional) GetMethod() string {
	return "taobao.tbk.dg.material.optional"
}

func (o *TaobaoTbkDgMaterialOptional) GetApiParams() Parameter {
	return o.params
}
