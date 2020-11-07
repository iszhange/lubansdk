package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-所有订单查询
// Link: https://open.taobao.com/api.htm?docId=40121&docType=2&scopeId=16175
// Author: Ken.Zhang
// Email: kenphp@yeah.net

type TbkRelationRefundResponse struct {
	Data      TbkRelationRefundData `json:"data"`
	RequestID string                `json:"request_id"`
}

type TbkRelationRefundData struct {
	HasNext       bool                         `json:"has_next"`
	HasPre        bool                         `json:"has_pre"`
	PageNo        int                          `json:"page_no"`
	PageSize      int                          `json:"page_size"`
	PositionIndex string                       `json:"position_index"`
	Results       []TbkRelationRefundOrderInfo `json:"results"`
}

type TbkRelationRefundOrderInfo struct {
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
	ItemImg                            string `json:"item_img"`
	ItemNum                            int    `json:"item_num"`
	ItemTitle                          string `json:"item_title"`
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
}

type TbkRelationRefund struct {
	params Parameter // 参数
}

func TbkRelationRefundRequest() *TbkRelationRefund {
	r := new(TbkRelationRefund)
	r.params = make(Parameter)
	return r
}

//func TbkRelationRefundResult(data []byte) (TbkRelationRefundData, error) {
//	var result TbkRelationRefundResponse
//	err := json.Unmarshal(data, &result)
//	if err != nil {
//		return TbkRelationRefundData{}, err
//	}
//	if result.RequestID == "" {
//		var errResult lubantop.ErrorResult
//		err = json.Unmarshal(data, &errResult)
//		return TbkRelationRefundData{}, errors.New(errResult.ErrorResponse.SubMsg)
//	}
//
//	return result.Data, nil
//}

func (r *TbkRelationRefund) SetBizType(value int) {
	r.params["biz_type"] = value
}

func (r *TbkRelationRefund) SetStartTime(value string) {
	r.params["start_time"] = value
}

func (r *TbkRelationRefund) SetRefundType(value int) {
	r.params["refund_type"] = value
}

func (r *TbkRelationRefund) SetSearchType(value int) {
	r.params["search_type"] = value
}

func (r *TbkRelationRefund) SetPageNo(value int) {
	r.params["page_no"] = value
}

func (r *TbkRelationRefund) SetPageSize(value int) {
	r.params["page_size"] = value
}

func (r *TbkRelationRefund) GetMethod() string {
	return "taobao.tbk.relation.refund"
}

func (r *TbkRelationRefund) GetApiParams() Parameter {
	p, _ := json.Marshal(r.params)
	var result = make(Parameter)
	result["search_option"] = string(p)

	return result
}
