package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-官方活动转链
// Link: https://open.taobao.com/api.htm?docId=48340&docType=2&scopeId=18294
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkActivityInfoGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkActivityInfoGet) New() *TaobaoTbkActivityInfoGet {
	r := new(TaobaoTbkActivityInfoGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkActivityInfoGet) Result(data []byte) (TaobaoTbkActivityInfoGetResponse, error) {
	var result TaobaoTbkActivityInfoGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkActivityInfoGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkActivityInfoGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkActivityInfoGet) SetAdzoneId(value int) {
	o.params["adzone_id"] = value
}

func (o *TaobaoTbkActivityInfoGet) SetSubPid(value string) {
	o.params["sub_pid"] = value
}

func (o *TaobaoTbkActivityInfoGet) SetRelationId(value int) {
	o.params["relation_id"] = value
}

func (o *TaobaoTbkActivityInfoGet) SetActivityMaterialId(value string) {
	o.params["activity_material_id"] = value
}

func (o *TaobaoTbkActivityInfoGet) SetUnionId(value string) {
	o.params["union_id"] = value
}

func (o *TaobaoTbkActivityInfoGet) GetMethod() string {
	return "taobao.tbk.activity.info.get"
}

func (o *TaobaoTbkActivityInfoGet) GetApiParams() Parameter {
	return o.params
}
