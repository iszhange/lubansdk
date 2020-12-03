package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-淘礼金发放及使用报表
// Link: https://open.taobao.com/api.htm?docId=43317&docType=2&scopeId=15029
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgVegasTljInstanceReport struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgVegasTljInstanceReport) New() *TaobaoTbkDgVegasTljInstanceReport {
	r := new(TaobaoTbkDgVegasTljInstanceReport)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgVegasTljInstanceReport) Result(data []byte) (TaobaoTbkDgVegasTljInstanceReportResponse, error) {
	var result TaobaoTbkDgVegasTljInstanceReportResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgVegasTljInstanceReportResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgVegasTljInstanceReportResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgVegasTljInstanceReport) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkDgVegasTljInstanceReport) GetMethod() string {
	return "taobao.tbk.dg.vegas.tlj.instance.report"
}

func (o *TaobaoTbkDgVegasTljInstanceReport) GetApiParams() Parameter {
	return o.params
}
