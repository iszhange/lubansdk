package requests

import (
	"encoding/json"
)

// 聚划算商品搜索接口
// Link: https://open.taobao.com/api.htm?docId=28762&docType=2&scopeId=16517
// Author: zhange
// Email: kenphp@yeah.net

type TaobaoJuItemsSearch struct {
	params Parameter
}

func (o TaobaoJuItemsSearch) New() *TaobaoJuItemsSearch {
	r := new(TaobaoJuItemsSearch)
	r.params = make(Parameter)
	return r
}

func (o TaobaoJuItemsSearch) Result(data []byte) (TaobaoJuItemsSearchResponse, error) {
	var result TaobaoJuItemsSearchResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoJuItemsSearchResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoJuItemsSearchResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoJuItemsSearch) SetCurrentPage(value int) {
	o.params["current_page"] = value
}

func (o *TaobaoJuItemsSearch) SetPageSize(value int) {
	o.params["page_size"] = value
}

func (o *TaobaoJuItemsSearch) SetPid(value string) {
	o.params["pid"] = value
}

func (o *TaobaoJuItemsSearch) SetPostage(value bool) {
	o.params["postage"] = value
}

func (o *TaobaoJuItemsSearch) SetStatus(value int) {
	o.params["status"] = value
}

func (o *TaobaoJuItemsSearch) SetTaobaoCategoryId(value int) {
	o.params["taobao_category_id"] = value
}

func (o *TaobaoJuItemsSearch) SetWord(value string) {
	o.params["word"] = value
}

func (o *TaobaoJuItemsSearch) GetMethod() string {
	return "taobao.ju.items.search"
}

func (o *TaobaoJuItemsSearch) GetApiParams() Parameter {
	p, _ := json.Marshal(o.params)
	var result = make(Parameter)
	result["param_top_item_query"] = string(p)

	return result
}
