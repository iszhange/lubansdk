package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-淘礼金创建
// Link: https://open.taobao.com/api.htm?docId=40173&docType=2&scopeId=15029
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgVegasTljCreate struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgVegasTljCreate) New() *TaobaoTbkDgVegasTljCreate {
	r := new(TaobaoTbkDgVegasTljCreate)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgVegasTljCreate) Result(data []byte) (TaobaoTbkDgVegasTljCreateResponse, error) {
	var result TaobaoTbkDgVegasTljCreateResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgVegasTljCreateResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgVegasTljCreateResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgVegasTljCreate) SetCampaignType(value string) {
	o.params["campaign_type"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetAdzoneId(value string) {
	o.params["adzone_id"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetItemId(value string) {
	o.params["item_id"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetTotalNum(value int) {
	o.params["total_num"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetName(value string) {
	o.params["name"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetUserTotalWinNumLimit(value int) {
	o.params["user_total_win_num_limit"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetSecuritySwitch(value bool) {
	o.params["security_switch"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetPerFace(value string) {
	o.params["per_face"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetSendStartTime(value string) {
	o.params["send_start_time"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetSendEndTime(value string) {
	o.params["send_end_time"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetUseEndTime(value string) {
	o.params["use_end_time"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetUseEndTimeMode(value int) {
	o.params["use_end_time_mode"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetUseStartTime(value string) {
	o.params["use_start_time"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) SetSecurityLevel(value int) {
	o.params["security_level"] = value
}

func (o *TaobaoTbkDgVegasTljCreate) GetMethod() string {
	return "taobao.tbk.dg.vegas.tlj.create"
}

func (o *TaobaoTbkDgVegasTljCreate) GetApiParams() Parameter {
	return o.params
}
