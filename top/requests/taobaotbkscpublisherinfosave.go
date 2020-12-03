package requests

import (
	"encoding/json"
)

// 淘宝客-公用-私域用户备案
// Link: https://open.taobao.com/api.htm?docId=37988&docType=2&scopeId=14474
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkScPublisherInfoSave struct {
	params Parameter // 参数
}

func (o TaobaoTbkScPublisherInfoSave) New() *TaobaoTbkScPublisherInfoSave {
	r := new(TaobaoTbkScPublisherInfoSave)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkScPublisherInfoSave) Result(data []byte) (TaobaoTbkScPublisherInfoSaveResponse, error) {
	var result TaobaoTbkScPublisherInfoSaveResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkScPublisherInfoSaveResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkScPublisherInfoSaveResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkScPublisherInfoSave) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkScPublisherInfoSave) GetMethod() string {
	return "taobao.tbk.sc.publisher.info.save"
}

func (o *TaobaoTbkScPublisherInfoSave) GetApiParams() Parameter {
	return o.params
}
