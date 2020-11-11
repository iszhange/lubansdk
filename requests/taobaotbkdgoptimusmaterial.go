package requests

import (
	"encoding/json"
)

// 淘宝客-公用-阿里妈妈推广券详情查询
// Link: https://open.taobao.com/api.htm?docId=31106&docType=2&scopeId=16189
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgOptimusMaterial struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgOptimusMaterial) New() *TaobaoTbkDgOptimusMaterial {
	r := new(TaobaoTbkDgOptimusMaterial)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgOptimusMaterial) Result(data []byte) (TaobaoTbkDgOptimusMaterialResponse, error) {
	var result TaobaoTbkDgOptimusMaterialResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgOptimusMaterialResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgOptimusMaterialResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgOptimusMaterial) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetPageNo(value int) {
	o.params["page_no"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetAdzoneId(value int) {
	o.params["adzone_id"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetMaterialId(value int) {
	o.params["material_id"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetDeviceValue(value string) {
	o.params["device_value"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetDeviceEncrypt(value string) {
	o.params["device_encrypt"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetDeviceType(value string) {
	o.params["device_type"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetContentId(value int) {
	o.params["content_id"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetContentSource(value string) {
	o.params["content_source"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetItemId(value int) {
	o.params["item_id"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) SetFavoritesId(value string) {
	o.params["favorites_id"] = value
}

func (o *TaobaoTbkDgOptimusMaterial) GetMethod() string {
	return "taobao.tbk.dg.optimus.material"
}

func (o *TaobaoTbkDgOptimusMaterial) GetApiParams() Parameter {
	return o.params
}
