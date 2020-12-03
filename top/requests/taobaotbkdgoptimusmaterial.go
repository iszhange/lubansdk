package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-物料精选
// Link: https://open.taobao.com/api.htm?docId=33947&docType=2
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgOptimusMaterial struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgOptimusMaterial) New() *TaobaoTbkDgOptimusMaterial {
	r := new(TaobaoTbkDgOptimusMaterial)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgOptimusMaterial) Result(data []byte) (TaobaoTbkDgOptimusMaterialResponse, error) {
	var result TaobaoTbkDgOptimusMaterialResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgOptimusMaterialResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgOptimusMaterialResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgOptimusMaterial) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkDgOptimusMaterial) GetMethod() string {
	return "taobao.tbk.dg.optimus.material"
}

func (o *TaobaoTbkDgOptimusMaterial) GetApiParams() Parameter {
	return o.params
}
