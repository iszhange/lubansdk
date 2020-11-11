package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-商品链接转换
// Link: https://open.taobao.com/api.htm?docId=24516&docType=2&scopeId=11653
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkItemConvert struct {
	params Parameter // 参数
}

func (o TaobaoTbkItemConvert) New() *TaobaoTbkItemConvert {
	r := new(TaobaoTbkItemConvert)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkItemConvert) Result(data []byte) (TaobaoTbkItemConvertResponse, error) {
	var result TaobaoTbkItemConvertResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkItemConvertResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkItemConvertResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkItemConvert) SetFields(value string) {
	o.params["fields"] = value
}

func (o *TaobaoTbkItemConvert) SetNumIids(value string) {
	o.params["num_iids"] = value
}

func (o *TaobaoTbkItemConvert) SetAdzoneId(value string) {
	o.params["adzone_id"] = value
}

func (o *TaobaoTbkItemConvert) SetPlatform(value int) {
	o.params["platform"] = value
}

func (o *TaobaoTbkItemConvert) SetUnid(value string) {
	o.params["unid"] = value
}

func (o *TaobaoTbkItemConvert) SetDx(value string) {
	o.params["dx"] = value
}

func (o *TaobaoTbkItemConvert) GetMethod() string {
	return "taobao.tbk.item.convert"
}

func (o *TaobaoTbkItemConvert) GetApiParams() Parameter {
	return o.params
}
