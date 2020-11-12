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

func (o *TaobaoTbkScPublisherInfoGet) SetInfoType(value int) {
	o.params["info_type"] = value
}

func (o *TaobaoTbkScPublisherInfoGet) SetRelationId(value string) {
	o.params["relation_id"] = value
}

func (o *TaobaoTbkScPublisherInfoGet) SetPageNo(value int) {
	o.params["page_no"] = value
}

func (o *TaobaoTbkScPublisherInfoGet) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkScPublisherInfoGet) SetRelationApp(value string) {
	o.params["relation_app"] = value
}

func (o *TaobaoTbkScPublisherInfoGet) SetSpecialId(value string) {
	o.params["special_id"] = value
}

func (o *TaobaoTbkScPublisherInfoGet) SetExternalId(value string) {
	o.params["external_id"] = value
}

func (o *TaobaoTbkScPublisherInfoGet) GetMethod() string {
	return "taobao.tbk.sc.publisher.info.get"
}

func (o *TaobaoTbkScPublisherInfoGet) GetApiParams() Parameter {
	return o.params
}
