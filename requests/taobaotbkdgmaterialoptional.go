package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-物料搜索
// Link: https://open.taobao.com/api.htm?docId=35896&docType=2&scopeId=16516
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoTbkDgMaterialOptional struct {
	params Parameter // 参数
}

func (o TaobaoTbkDgMaterialOptional) New() *TaobaoTbkDgMaterialOptional {
	r := new(TaobaoTbkDgMaterialOptional)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkDgMaterialOptional) Result(data []byte) (TaobaoTbkDgMaterialOptionalResponse, error) {
	var result TaobaoTbkDgMaterialOptionalResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkDgMaterialOptionalResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkDgMaterialOptionalResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkDgMaterialOptional) SetStartDsr(value int) {
	o.params["start_dsr"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetPageNo(value int) {
	o.params["page_no"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetPlatform(value int) {
	o.params["platform"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetEndTkRate(value int) {
	o.params["end_tk_rate"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetStartTkRate(value int) {
	o.params["start_tk_rate"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetEndPrice(value int) {
	o.params["end_price"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetStartPrice(value int) {
	o.params["start_price"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetIsOverseas(value bool) {
	o.params["is_overseas"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetIsTmall(value bool) {
	o.params["is_tmall"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetSort(value string) {
	o.params["sort"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetItemloc(value string) {
	o.params["itemloc"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetCat(value string) {
	o.params["cat"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetQ(value string) {
	o.params["q"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetMaterialId(value int) {
	o.params["material_id"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetHasCoupon(value bool) {
	o.params["has_coupon"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetIp(value string) {
	o.params["ip"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetAdzoneId(value string) {
	o.params["adzone_id"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetNeedFreeShipment(value bool) {
	o.params["need_free_shipment"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetNeedPrepay(value bool) {
	o.params["need_prepay"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetIncludePayRate30(value bool) {
	o.params["include_pay_rate_30"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetIncludeGoodRate(value bool) {
	o.params["include_good_rate"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetIncludeRfdRate(value bool) {
	o.params["include_rfd_rate"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetNpxLevel(value int) {
	o.params["npx_level"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetEndKaTkRate(value int) {
	o.params["end_ka_tk_rate"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetStartKaTkRate(value int) {
	o.params["start_ka_tk_rate"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetDeviceEncrypt(value string) {
	o.params["device_encrypt"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetDeviceValue(value string) {
	o.params["device_value"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetDeviceType(value string) {
	o.params["device_type"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetLockRateEndTime(value int) {
	o.params["lock_rate_end_time"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetLockRateStartTime(value int) {
	o.params["lock_rate_start_time"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetLongitude(value string) {
	o.params["longitude"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetLatitude(value string) {
	o.params["latitude"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetCityCode(value string) {
	o.params["city_code"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetSellerIds(value string) {
	o.params["seller_ids"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetSpecialId(value string) {
	o.params["special_id"] = value
}

func (o *TaobaoTbkDgMaterialOptional) SetRelationId(value string) {
	o.params["relation_id"] = value
}

func (o *TaobaoTbkDgMaterialOptional) GetMethod() string {
	return "taobao.tbk.dg.material.optional"
}

func (o *TaobaoTbkDgMaterialOptional) GetApiParams() Parameter {
	return o.params
}
