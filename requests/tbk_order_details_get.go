package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-所有订单查询
// Link: https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
// Author: Ken.Zhang
// Email: kenphp@yeah.net

type TbkOrderDetailsGetResponse struct {
	Data      TbkOrderDetailsGetData `json:"data"`
	RequestID string                 `json:"request_id"`
}

type TbkOrderDetailsGetData struct {
	HasNext       bool                          `json:"has_next"`
	HasPre        bool                          `json:"has_pre"`
	PageNo        int                           `json:"page_no"`
	PageSize      int                           `json:"page_size"`
	PositionIndex string                        `json:"position_index"`
	Results       []TbkOrderDetailsGetOrderInfo `json:"results"`
}

type TbkOrderDetailsGetOrderInfo struct {
	AdzoneID                           int64  `json:"adzone_id"`
	AdzoneName                         string `json:"adzone_name"`
	AlimamaRate                        string `json:"alimama_rate"`
	AlimamaShareFee                    string `json:"alimama_share_fee"`
	AlipayTotalPrice                   string `json:"alipay_total_price"`
	ClickTime                          string `json:"click_time"`
	DepositPrice                       string `json:"deposit_price"`
	FlowSource                         string `json:"flow_source"`
	IncomeRate                         string `json:"income_rate"`
	ItemCategoryName                   string `json:"item_category_name"`
	ItemId                             int64  `json:"item_id"`
	ItemLink                           string `json:"item_link"`
	ItemImg                            string `json:"item_img"`
	ItemNum                            int    `json:"item_num"`
	ItemTitle                          string `json:"item_title"`
	ItemPrice                          string `json:"item_price"`
	OrderType                          string `json:"order_type"`
	PayPrice                           string `json:"pay_price"`
	PubID                              int    `json:"pub_id"`
	PubShareFee                        string `json:"pub_share_fee"`
	PubSharePreFee                     string `json:"pub_share_pre_fee"`
	PubShareRate                       string `json:"pub_share_rate"`
	RefundTag                          int    `json:"refund_tag"`
	RelationID                         int    `json:"relation_id"`
	SellerNick                         string `json:"seller_nick"`
	SellerShopTitle                    string `json:"seller_shop_title"`
	SiteID                             int    `json:"site_id"`
	SiteName                           string `json:"site_name"`
	SubsidyFee                         string `json:"subsidy_fee"`
	SubsidyRate                        string `json:"subsidy_rate"`
	SubsidyType                        string `json:"subsidy_type"`
	TbDepositTime                      string `json:"tb_deposit_time"`
	TbPaidTime                         string `json:"tb_paid_time"`
	TerminalType                       string `json:"terminal_type"`
	TkCommissionFeeForMediaPlatform    string `json:"tk_commission_fee_for_media_platform"`
	TkCommissionPreFeeForMediaPlatform string `json:"tk_commission_pre_fee_for_media_platform"`
	TkCommissionRateForMediaPlatform   string `json:"tk_commission_rate_for_media_platform"`
	TkCreateTime                       string `json:"tk_create_time"`
	TkDepositTime                      string `json:"tk_deposit_time"`
	TkEarningTime                      string `json:"tk_earning_time"`
	TkOrderRole                        int    `json:"tk_order_role"`
	TkPaidTime                         string `json:"tk_paid_time"`
	TkStatus                           int    `json:"tk_status"`
	TkTotalRate                        string `json:"tk_total_rate"`
	TotalCommissionFee                 string `json:"total_commission_fee"`
	TotalCommissionRate                string `json:"total_commission_rate"`
	TradeID                            string `json:"trade_id"`
	TradeParentID                      string `json:"trade_parent_id"`
	ServiceFeeDtoList                  []struct {
		ShareFee          string `json:"share_fee"`
		SharePreFee       string `json:"share_pre_fee"`
		ShareRelativeRate string `json:"share_relative_rate"`
		TkShareRoleType   int    `json:"tk_share_role_type"`
	} `json:"service_fee_dto_list"`
}

type TbkOrderDetailsGet struct {
	params Parameter // 参数
}

func TbkOrderDetailsGetRequest() *TbkOrderDetailsGet {
	r := new(TbkOrderDetailsGet)
	r.params = make(Parameter)
	return r
}

func TbkOrderDetailsGetResult(data []byte) (TbkOrderDetailsGetData, error) {
	var result TbkOrderDetailsGetResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TbkOrderDetailsGetData{}, err
	}
	if result.RequestID == "" {
		return TbkOrderDetailsGetData{}, API_RESPONSE_ERROR
	}

	return result.Data, nil
}

//func TbkOrderDetailGetError(data []byte) (lubantop.Error, error) {
//	var errResult lubantop.ErrorResult
//	err := json.Unmarshal(data, &errResult)
//	if err != nil {
//		return lubantop.Error{}, err
//	}
//	return errResult.ErrorResponse, nil
//}

func (r *TbkOrderDetailsGet) SetQueryType(value int) {
	r.params["query_type"] = value
}

func (r *TbkOrderDetailsGet) SetOrderScene(value int) {
	r.params["order_scene"] = value
}

func (r *TbkOrderDetailsGet) SetPositionIndex(value string) {
	r.params["position_index"] = value
}

func (r *TbkOrderDetailsGet) SetMemberType(value int) {
	r.params["member_type"] = value
}

func (r *TbkOrderDetailsGet) SetTkStatus(value int) {
	r.params["tk_status"] = value
}

func (r *TbkOrderDetailsGet) SetStartTime(value string) {
	r.params["start_time"] = value
}

func (r *TbkOrderDetailsGet) SetEndTime(value string) {
	r.params["end_time"] = value
}

func (r *TbkOrderDetailsGet) SetJumpType(value int) {
	r.params["jump_type"] = value
}

func (r *TbkOrderDetailsGet) SetPageNo(value int) {
	r.params["page_no"] = value
}

func (r *TbkOrderDetailsGet) SetPageSize(value int) {
	r.params["page_size"] = value
}

func (r *TbkOrderDetailsGet) GetMethod() string {
	return "taobao.tbk.order.details.get"
}

func (r *TbkOrderDetailsGet) GetApiParams() Parameter {
	return r.params
}
