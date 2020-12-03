package requests

import (
	"encoding/json"
)

// 淘宝客-服务商-查询超级红包发放个数
// Link: https://open.taobao.com/api.htm?docId=47590&docType=2&scopeId=17875
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkScVegasSendReport struct {
	params Parameter // 参数
}

func (o TaobaoTbkScVegasSendReport) New() *TaobaoTbkScVegasSendReport {
	r := new(TaobaoTbkScVegasSendReport)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkScVegasSendReport) Result(data []byte) (TaobaoTbkScVegasSendReportResponse, error) {
	var result TaobaoTbkScVegasSendReportResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkScVegasSendReportResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkScVegasSendReportResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkScVegasSendReport) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkScVegasSendReport) GetMethod() string {
	return "taobao.tbk.sc.vegas.send.report"
}

func (o *TaobaoTbkScVegasSendReport) GetApiParams() Parameter {
	return o.params
}
