package requests

import (
	"encoding/json"
)

// 淘宝客-公用-淘口令生成
// Link: https://open.taobao.com/api.htm?docId=31127&docType=2&scopeId=11655
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkTpwdCreate struct {
	params Parameter // 参数
}

func (o TaobaoTbkTpwdCreate) New() *TaobaoTbkTpwdCreate {
	r := new(TaobaoTbkTpwdCreate)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkTpwdCreate) Result(data []byte) (TaobaoTbkTpwdCreateResponse, error) {
	var result TaobaoTbkTpwdCreateResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkTpwdCreateResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkTpwdCreateResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkTpwdCreate) SetUserId(value string) {
	o.params["user_id"] = value
}

func (o *TaobaoTbkTpwdCreate) SetText(value string) {
	o.params["text"] = value
}

func (o *TaobaoTbkTpwdCreate) SetUrl(value string) {
	o.params["url"] = value
}

func (o *TaobaoTbkTpwdCreate) SetLogo(value bool) {
	o.params["logo"] = value
}

func (o *TaobaoTbkTpwdCreate) GetMethod() string {
	return "taobao.tbk.tpwd.create"
}

func (o *TaobaoTbkTpwdCreate) GetApiParams() Parameter {
	return o.params
}
