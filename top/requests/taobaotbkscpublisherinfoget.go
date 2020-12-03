package requests

import (
	"encoding/json"
)

// 淘宝客-公用-私域用户备案信息查询
// Link: https://open.taobao.com/api.htm?docId=37989&docType=2&scopeId=14474
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkScPublisherInfoGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkScPublisherInfoGet) New() *TaobaoTbkScPublisherInfoGet {
	r := new(TaobaoTbkScPublisherInfoGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkScPublisherInfoGet) Result(data []byte) (TaobaoTbkScPublisherInfoGetResponse, error) {
	var result TaobaoTbkScPublisherInfoGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkScPublisherInfoGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkScPublisherInfoGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkScPublisherInfoGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkScPublisherInfoGet) GetMethod() string {
	return "taobao.tbk.sc.publisher.info.get"
}

func (o *TaobaoTbkScPublisherInfoGet) GetApiParams() Parameter {
	return o.params
}
