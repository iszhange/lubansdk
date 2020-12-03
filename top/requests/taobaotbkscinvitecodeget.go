package requests

import (
	"encoding/json"
)

// 淘宝客-公用-私域用户邀请码生成
// Link: https://open.taobao.com/api.htm?docId=38046&docType=2&scopeId=14474
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkScInvitecodeGet struct {
	params Parameter // 参数
}

func (o TaobaoTbkScInvitecodeGet) New() *TaobaoTbkScInvitecodeGet {
	r := new(TaobaoTbkScInvitecodeGet)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkScInvitecodeGet) Result(data []byte) (TaobaoTbkScInvitecodeGetResponse, error) {
	var result TaobaoTbkScInvitecodeGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkScInvitecodeGetResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkScInvitecodeGetResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkScInvitecodeGet) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkScInvitecodeGet) GetMethod() string {
	return "taobao.tbk.sc.invitecode.get"
}

func (o *TaobaoTbkScInvitecodeGet) GetApiParams() Parameter {
	return o.params
}
