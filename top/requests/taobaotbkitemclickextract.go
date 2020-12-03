package requests

import (
	"encoding/json"
)

// 淘宝客-公用-链接解析出商品id
// Link: https://open.taobao.com/API.htm?docId=28156&docType=2
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkItemClickExtract struct {
	params Parameter // 参数
}

func (o TaobaoTbkItemClickExtract) New() *TaobaoTbkItemClickExtract {
	r := new(TaobaoTbkItemClickExtract)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkItemClickExtract) Result(data []byte) (TaobaoTbkItemClickExtractResponse, error) {
	var result TaobaoTbkItemClickExtractResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkItemClickExtractResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkItemClickExtractResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkItemClickExtract) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkItemClickExtract) GetMethod() string {
	return "taobao.tbk.item.click.extract"
}

func (o *TaobaoTbkItemClickExtract) GetApiParams() Parameter {
	return o.params
}
