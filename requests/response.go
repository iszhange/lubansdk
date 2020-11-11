package requests

import "errors"

var (
	NIL_TYPE_ERROR     = errors.New("数据类型为nil")
	UNKONWN_TYPE_ERROR = errors.New("未知的数据类型")
	API_RESPONSE_ERROR = errors.New("接口返回错误")
	EMPTY_RESPONSE     = errors.New("响应为空")
)

type ErrorResponse struct {
	ErrorResponse struct {
		Code      int    `json:"code"`
		Msg       string `json:"msg"`
		SubCode   string `json:"sub_code"`
		SubMsg    string `json:"sub_msg"`
		RequestID string `json:"request_id"`
	} `json:"error_response"`
}

type TaobaoTbkItemInfoGetResponse struct {
	Results []struct {
		CatLeafName                string   `json:"cat_leaf_name"`
		CatName                    string   `json:"cat_name"`
		FreeShipment               bool     `json:"free_shipment"`
		ItemURL                    string   `json:"item_url"`
		JuOnlineEndTime            string   `json:"ju_online_end_time"`
		JuOnlineStartTime          string   `json:"ju_online_start_time"`
		JuPreShowEndTime           string   `json:"ju_pre_show_end_time"`
		JuPreShowStartTime         string   `json:"ju_pre_show_start_time"`
		KuadianPromotionInfo       string   `json:"kuadian_promotion_info"`
		MaterialLibType            string   `json:"material_lib_type"`
		Nick                       string   `json:"nick"`
		NumIid                     int64    `json:"num_iid"`
		PictURL                    string   `json:"pict_url"`
		PresaleDeposit             string   `json:"presale_deposit"`
		PresaleEndTime             int      `json:"presale_end_time"`
		PresaleStartTime           int      `json:"presale_start_time"`
		PresaleTailEndTime         int      `json:"presale_tail_end_time"`
		PresaleTailStartTime       int      `json:"presale_tail_start_time"`
		Provcity                   string   `json:"provcity"`
		ReservePrice               string   `json:"reserve_price"`
		SalePrice                  string   `json:"sale_price"`
		SellerID                   int64    `json:"seller_id"`
		SmallImages                []string `json:"small_images"`
		SuperiorBrand              string   `json:"superior_brand"`
		Title                      string   `json:"title"`
		TmallPlayActivityEndTime   int      `json:"tmall_play_activity_end_time"`
		TmallPlayActivityStartTime int      `json:"tmall_play_activity_start_time"`
		UserType                   int      `json:"user_type"`
		Volume                     int      `json:"volume"`
		ZkFinalPrice               string   `json:"zk_final_price"`
	} `json:"results"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkOrderDetailsGetResponse struct {
	Data struct {
		HasNext       bool   `json:"has_next"`
		HasPre        bool   `json:"has_pre"`
		PageNo        int    `json:"page_no"`
		PageSize      int    `json:"page_size"`
		PositionIndex string `json:"position_index"`
		Results       []struct {
			AdzoneID          int64  `json:"adzone_id"`
			AdzoneName        string `json:"adzone_name"`
			AlimamaRate       string `json:"alimama_rate"`
			AlimamaShareFee   string `json:"alimama_share_fee"`
			AlipayTotalPrice  string `json:"alipay_total_price,omitempty"`
			ClickTime         string `json:"click_time"`
			DepositPrice      string `json:"deposit_price"`
			FlowSource        string `json:"flow_source"`
			IncomeRate        string `json:"income_rate"`
			IsLx              string `json:"is_lx"`
			ItemCategoryName  string `json:"item_category_name"`
			ItemID            int64  `json:"item_id"`
			ItemImg           string `json:"item_img"`
			ItemLink          string `json:"item_link"`
			ItemNum           int    `json:"item_num"`
			ItemPrice         string `json:"item_price"`
			ItemTitle         string `json:"item_title"`
			OrderType         string `json:"order_type"`
			PayPrice          string `json:"pay_price"`
			PubID             int    `json:"pub_id"`
			PubShareFee       string `json:"pub_share_fee"`
			PubSharePreFee    string `json:"pub_share_pre_fee"`
			PubShareRate      string `json:"pub_share_rate"`
			RefundTag         int    `json:"refund_tag"`
			RelationID        int    `json:"relation_id"`
			SellerNick        string `json:"seller_nick"`
			SellerShopTitle   string `json:"seller_shop_title"`
			ServiceFeeDtoList []struct {
				ShareFee          string `json:"share_fee"`
				SharePreFee       string `json:"share_pre_fee"`
				ShareRelativeRate string `json:"share_relative_rate"`
				TkShareRoleType   int    `json:"tk_share_role_type"`
			} `json:"service_fee_dto_list"`
			SiteID                             int    `json:"site_id"`
			SiteName                           string `json:"site_name"`
			SubsidyFee                         string `json:"subsidy_fee"`
			SubsidyRate                        string `json:"subsidy_rate"`
			SubsidyType                        string `json:"subsidy_type"`
			TbDepositTime                      string `json:"tb_deposit_time"`
			TbPaidTime                         string `json:"tb_paid_time,omitempty"`
			TerminalType                       string `json:"terminal_type"`
			TkCommissionFeeForMediaPlatform    string `json:"tk_commission_fee_for_media_platform"`
			TkCommissionPreFeeForMediaPlatform string `json:"tk_commission_pre_fee_for_media_platform"`
			TkCommissionRateForMediaPlatform   string `json:"tk_commission_rate_for_media_platform"`
			TkCreateTime                       string `json:"tk_create_time"`
			TkDepositTime                      string `json:"tk_deposit_time"`
			TkOrderRole                        int    `json:"tk_order_role"`
			TkPaidTime                         string `json:"tk_paid_time,omitempty"`
			TkEarningTime                      string `json:"tk_earning_time"`
			TkStatus                           int    `json:"tk_status"`
			TkTotalRate                        string `json:"tk_total_rate"`
			TotalCommissionFee                 string `json:"total_commission_fee"`
			TotalCommissionRate                string `json:"total_commission_rate"`
			TradeID                            string `json:"trade_id"`
			TradeParentID                      string `json:"trade_parent_id"`
		} `json:"results"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkRelationRefundResponse struct {
	Result struct {
		BizErrorCode int `json:"biz_error_code"`
		Data         struct {
			PageNo   string `json:"page_no"`
			PageSize string `json:"page_size"`
			Results  []struct {
				EarningTime              string `json:"earning_time"`
				RefundFee                string `json:"refund_fee"`
				RefundStatus             int    `json:"refund_status"`
				RefundType               int    `json:"refund_type"`
				RelationID               int    `json:"relation_id"`
				TbAuctionTitle           string `json:"tb_auction_title"`
				TbTradeCreateTime        string `json:"tb_trade_create_time"`
				TbTradeFinishPrice       string `json:"tb_trade_finish_price"`
				TbTradeID                int64  `json:"tb_trade_id"`
				TbTradeParentID          int64  `json:"tb_trade_parent_id"`
				TkCommissionFeeRefundPub string `json:"tk_commission_fee_refund_pub"`
				TkPubID                  int    `json:"tk_pub_id"`
				TkPubShowReturnFee       string `json:"tk_pub_show_return_fee"`
				TkRefundSuitTime         string `json:"tk_refund_suit_time"`
				TkRefundTime             string `json:"tk_refund_time"`
				TkSubsidyFeeRefundPub    string `json:"tk_subsidy_fee_refund_pub"`
			} `json:"results"`
			TotalCount string `json:"total_count"`
		} `json:"data"`
		ResultCode int `json:"result_code"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkDgVegasSendStatusResponse struct {
	Data struct {
		ResultList []struct {
			IsNewUser string `json:"is_new_user"`
		} `json:"result_list"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkDgVegasSendReportResponse struct {
	Result struct {
		Model struct {
			RelationRptList []struct {
				BizDate    string `json:"biz_date"`
				FundNum    int    `json:"fund_num"`
				RelationID int    `json:"relation_id"`
			} `json:"relation_rpt_list"`
		} `json:"model"`
		Success bool `json:"success"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkScVegasSendReportResponse struct {
	Result struct {
		Model struct {
			RelationRptList []struct {
				BizDate    string `json:"biz_date"`
				FundNum    int    `json:"fund_num"`
				RelationID int    `json:"relation_id"`
			} `json:"relation_rpt_list"`
		} `json:"model"`
		Success bool `json:"success"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkSpreadGetResponse struct {
	Results []struct {
		Content string `json:"content"`
		ErrMsg  string `json:"err_msg"`
	} `json:"results"`
	TotalResults int    `json:"total_results"`
	RequestID    string `json:"request_id"`
}

type TaobaoTbkActivityInfoGetResponse struct {
	Data struct {
		ClickURL       string `json:"click_url"`
		MaterialOssURL string `json:"material_oss_url"`
		PageEndTime    string `json:"page_end_time"`
		PageName       string `json:"page_name"`
		PageStartTime  string `json:"page_start_time"`
		TerminalType   string `json:"terminal_type"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkDgVegasTljCreateResponse struct {
	Result struct {
		Model struct {
			AvailableFee string `json:"available_fee"`
			RightsID     string `json:"rights_id"`
			SendURL      string `json:"send_url"`
			VegasCode    string `json:"vegas_code"`
		} `json:"model"`
		Success bool `json:"success"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkDgVegasTljInstanceReportResponse struct {
	Result struct {
		Model struct {
			UnfreezeAmount      string `json:"unfreeze_amount"`
			UnfreezeNum         int    `json:"unfreeze_num"`
			RefundAmount        string `json:"refund_amount"`
			RefundNum           int    `json:"refund_num"`
			AlipayAmount        string `json:"alipay_amount"`
			UseAmount           string `json:"use_amount"`
			UseNum              int    `json:"use_num"`
			WinAmount           string `json:"win_amount"`
			WinNum              int    `json:"win_num"`
			PreCommissionAmount string `json:"pre_commission_amount"`
			FpRefundAmount      string `json:"fp_refund_amount"`
			FpRefundNum         int    `json:"fp_refund_num"`
		} `json:"model"`
		Success bool `json:"success"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type TaobaoJuItemsSearchResponse struct {
	Result struct {
		CurrentPage int `json:"current_page"`
		ModelList   []struct {
			ActPrice        string   `json:"act_price"`
			CategoryName    string   `json:"category_name"`
			ItemID          int64    `json:"item_id"`
			ItemUspList     []string `json:"item_usp_list"`
			JuID            int64    `json:"ju_id"`
			OnlineEndTime   int64    `json:"online_end_time"`
			OnlineStartTime int64    `json:"online_start_time"`
			OrigPrice       string   `json:"orig_price"`
			PayPostage      bool     `json:"pay_postage"`
			PcURL           string   `json:"pc_url"`
			PicURLForPC     string   `json:"pic_url_for_p_c"`
			PicURLForWL     string   `json:"pic_url_for_w_l"`
			PlatformID      int      `json:"platform_id"`
			PriceUspList    []string `json:"price_usp_list"`
			ShowEndTime     int64    `json:"show_end_time"`
			ShowStartTime   int64    `json:"show_start_time"`
			TbFirstCatID    int      `json:"tb_first_cat_id"`
			Title           string   `json:"title"`
			UspDescList     []string `json:"usp_desc_list"`
			WapURL          string   `json:"wap_url"`
		} `json:"model_list"`
		PageSize    int  `json:"page_size"`
		Success     bool `json:"success"`
		TotalItem   int  `json:"total_item"`
		TotalPage   int  `json:"total_page"`
		TrackParams struct {
		} `json:"track_params"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkCouponGetResponse struct {
	Data struct {
		CouponActivityID  string `json:"coupon_activity_id"`
		CouponAmount      string `json:"coupon_amount"`
		CouponEndTime     string `json:"coupon_end_time"`
		CouponRemainCount int    `json:"coupon_remain_count"`
		CouponSrcScene    int    `json:"coupon_src_scene"`
		CouponStartFee    string `json:"coupon_start_fee"`
		CouponStartTime   string `json:"coupon_start_time"`
		CouponTotalCount  int    `json:"coupon_total_count"`
		CouponType        int    `json:"coupon_type"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}
