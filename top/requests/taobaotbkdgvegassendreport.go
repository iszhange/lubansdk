package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-查询超级红包发放个数
// Link: https://open.taobao.com/api.htm?docId=47593&docType=2&scopeId=17873
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgVegasSendReport struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgVegasSendReport) New() *TaobaoTbkDgVegasSendReport {
	r := new(TaobaoTbkDgVegasSendReport)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgVegasSendReport) Result(data []byte) (TaobaoTbkDgVegasSendReportResponse, error) {
	var result TaobaoTbkDgVegasSendReportResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgVegasSendReportResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgVegasSendReportResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgVegasSendReport) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkDgVegasSendReport) GetMethod() string {
	return "taobao.tbk.dg.vegas.send.report"
}

func (o *TaobaoTbkDgVegasSendReport) GetApiParams() Parameter {
	return o.params
}
