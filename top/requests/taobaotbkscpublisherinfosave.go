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

func (o *TaobaoTbkScPublisherInfoSave) SetRelationFrom(value string) {
	o.params["relation_from"] = value
}

func (o *TaobaoTbkScPublisherInfoSave) SetOfflineScene(value string) {
	o.params["offline_scene"] = value
}

func (o *TaobaoTbkScPublisherInfoSave) SetOnlineScene(value string) {
	o.params["online_scene"] = value
}

func (o *TaobaoTbkScPublisherInfoSave) SetInviterCode(value string) {
	o.params["inviter_code"] = value
}

func (o *TaobaoTbkScPublisherInfoSave) SetInfoType(value int) {
	o.params["info_type"] = value
}

func (o *TaobaoTbkScPublisherInfoSave) SetNote(value string) {
	o.params["note"] = value
}

func (o *TaobaoTbkScPublisherInfoSave) SetRegisterInfo(value string) {
	o.params["register_info"] = value
}

func (o *TaobaoTbkScPublisherInfoSave) GetMethod() string {
	return "taobao.tbk.sc.publisher.info.save"
}

func (o *TaobaoTbkScPublisherInfoSave) GetApiParams() Parameter {
	return o.params
}
