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
			Unid                               string `json:"unid"`
			LxRid                              string `json:"lx_rid"`
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

type TaobaoTbkShopRecommendGetResponse struct {
	Results []struct {
		PictURL    string `json:"pict_url"`
		SellerNick string `json:"seller_nick"`
		ShopTitle  string `json:"shop_title"`
		ShopType   string `json:"shop_type"`
		ShopURL    string `json:"shop_url"`
		UserID     int64  `json:"user_id"`
	} `json:"results"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkDgOptimusMaterialResponse struct {
	IsDefault  string `json:"is_default"`
	ResultList []struct {
		CategoryID                 int      `json:"category_id"`
		CategoryName               string   `json:"category_name"`
		ClickURL                   string   `json:"click_url"`
		CouponAmount               int      `json:"coupon_amount"`
		CouponEndTime              string   `json:"coupon_end_time"`
		CouponStartFee             string   `json:"coupon_start_fee"`
		CouponStartTime            string   `json:"coupon_start_time"`
		CouponTotalCount           int      `json:"coupon_total_count"`
		ItemDescription            string   `json:"item_description"`
		ItemID                     int64    `json:"item_id"`
		LevelOneCategoryID         int      `json:"level_one_category_id"`
		LevelOneCategoryName       string   `json:"level_one_category_name"`
		PictURL                    string   `json:"pict_url"`
		ReservePrice               string   `json:"reserve_price"`
		SellerID                   int64    `json:"seller_id"`
		ShopTitle                  string   `json:"shop_title"`
		ShortTitle                 string   `json:"short_title"`
		SmallImages                []string `json:"small_images"`
		SubTitle                   string   `json:"sub_title"`
		Title                      string   `json:"title"`
		TqgOnlineEndTime           string   `json:"tqg_online_end_time"`
		TqgOnlineStartTime         string   `json:"tqg_online_start_time"`
		TqgSoldCount               int      `json:"tqg_sold_count"`
		TqgTotalCount              int      `json:"tqg_total_count"`
		UserType                   int      `json:"user_type"`
		Volume                     int      `json:"volume"`
		WhiteImage                 string   `json:"white_image"`
		ZkFinalPrice               string   `json:"zk_final_price"`
		CommissionRate             string   `json:"commission_rate"`
		CouponClickURL             string   `json:"coupon_click_url"`
		CouponRemainCount          int      `json:"coupon_remain_count"`
		CouponShareURL             string   `json:"coupon_share_url"`
		JhsPriceUspList            string   `json:"jhs_price_usp_list"`
		JuPlayEndTime              int      `json:"ju_play_end_time"`
		JuPlayStartTime            int      `json:"ju_play_start_time"`
		Nick                       string   `json:"nick"`
		JuOnlineEndTime            string   `json:"ju_online_end_time"`
		JuOnlineStartTime          string   `json:"ju_online_start_time"`
		MaochaoPlayConditions      string   `json:"maochao_play_conditions"`
		MaochaoPlayDiscountType    string   `json:"maochao_play_discount_type"`
		MaochaoPlayDiscounts       string   `json:"maochao_play_discounts"`
		MaochaoPlayEndTime         string   `json:"maochao_play_end_time"`
		MaochaoPlayFreePostFee     string   `json:"maochao_play_free_post_fee"`
		MaochaoPlayStartTime       string   `json:"maochao_play_start_time"`
		TmallPlayActivityEndTime   int      `json:"tmall_play_activity_end_time"`
		TmallPlayActivityStartTime int      `json:"tmall_play_activity_start_time"`
		FavoritesInfo              struct {
			FavoritesList []struct {
				FavoritesID    int    `json:"favorites_id"`
				FavoritesTitle string `json:"favorites_title"`
			} `json:"favorites_list"`
			TotalCount int `json:"total_count"`
		} `json:"favorites_info"`
		XID                  string `json:"x_id"`
		PresaleDeposit       string `json:"presale_deposit"`
		PresaleEndTime       int    `json:"presale_end_time"`
		PresaleStartTime     int    `json:"presale_start_time"`
		PresaleTailEndTime   int    `json:"presale_tail_end_time"`
		PresaleTailStartTime int    `json:"presale_tail_start_time"`
	} `json:"result_list"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkDgOptimusPromotionResponse struct {
	ResultList []struct {
		ConditionType    string `json:"condition_type"`
		DiscountType     string `json:"discount_type"`
		DisplayEndTime   string `json:"display_end_time"`
		DisplayStartTime string `json:"display_start_time"`
		Nick             string `json:"nick"`
		PromotionExtend  struct {
			PromotionURL      string `json:"promotion_url"`
			RecommendItemList []struct {
				ItemID int64  `json:"item_id"`
				URL    string `json:"url"`
			} `json:"recommend_item_list"`
			YoujiaCouponInfo struct {
				ItemID string `json:"item_id"`
				URL    string `json:"url"`
			} `json:"youjia_coupon_info"`
		} `json:"promotion_extend"`
		PromotionList []struct {
			EntryCondition     string `json:"entry_condition"`
			EntryDiscount      string `json:"entry_discount"`
			EntryUsedEndTime   string `json:"entry_used_end_time"`
			EntryUsedStartTime string `json:"entry_used_start_time"`
		} `json:"promotion_list"`
		PromotionType  string `json:"promotion_type"`
		RemainCount    int    `json:"remain_count"`
		SellerID       string `json:"seller_id"`
		ShopPictureURL string `json:"shop_picture_url"`
		ShopTitle      string `json:"shop_title"`
		TotalCount     int    `json:"total_count"`
	} `json:"result_list"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkShopGetResponse struct {
	Results []struct {
		PictURL    string `json:"pict_url"`
		SellerNick string `json:"seller_nick"`
		ShopTitle  string `json:"shop_title"`
		ShopType   string `json:"shop_type"`
		ShopURL    string `json:"shop_url"`
		UserID     int    `json:"user_id"`
	} `json:"results"`
	TotalResults int    `json:"total_results"`
	RequestID    string `json:"request_id"`
}

type TaobaoTbkDgMaterialOptionalResponse struct {
	ResultList []struct {
		CategoryID           int      `json:"category_id"`
		CategoryName         string   `json:"category_name"`
		CommissionRate       string   `json:"commission_rate"`
		CommissionType       string   `json:"commission_type"`
		CouponAmount         string   `json:"coupon_amount,omitempty"`
		CouponEndTime        string   `json:"coupon_end_time,omitempty"`
		CouponID             string   `json:"coupon_id"`
		CouponInfo           string   `json:"coupon_info"`
		CouponRemainCount    int      `json:"coupon_remain_count"`
		CouponShareURL       string   `json:"coupon_share_url,omitempty"`
		CouponStartFee       string   `json:"coupon_start_fee,omitempty"`
		CouponStartTime      string   `json:"coupon_start_time,omitempty"`
		CouponTotalCount     int      `json:"coupon_total_count"`
		IncludeDxjh          string   `json:"include_dxjh"`
		IncludeMkt           string   `json:"include_mkt"`
		InfoDxjh             string   `json:"info_dxjh"`
		ItemDescription      string   `json:"item_description"`
		ItemID               int64    `json:"item_id"`
		ItemURL              string   `json:"item_url"`
		KuadianPromotionInfo string   `json:"kuadian_promotion_info,omitempty"`
		LevelOneCategoryID   int      `json:"level_one_category_id"`
		LevelOneCategoryName string   `json:"level_one_category_name"`
		Nick                 string   `json:"nick"`
		NumIid               int64    `json:"num_iid"`
		PictURL              string   `json:"pict_url"`
		PresaleDeposit       string   `json:"presale_deposit"`
		PresaleEndTime       int      `json:"presale_end_time"`
		PresaleStartTime     int      `json:"presale_start_time"`
		PresaleTailEndTime   int      `json:"presale_tail_end_time"`
		PresaleTailStartTime int      `json:"presale_tail_start_time"`
		Provcity             string   `json:"provcity"`
		RealPostFee          string   `json:"real_post_fee"`
		ReservePrice         string   `json:"reserve_price"`
		SellerID             int      `json:"seller_id"`
		ShopDsr              int      `json:"shop_dsr"`
		ShopTitle            string   `json:"shop_title"`
		ShortTitle           string   `json:"short_title"`
		SmallImages          []string `json:"small_images"`
		Title                string   `json:"title"`
		TkTotalCommi         string   `json:"tk_total_commi"`
		TkTotalSales         string   `json:"tk_total_sales"`
		URL                  string   `json:"url"`
		UserType             int      `json:"user_type"`
		Volume               int      `json:"volume"`
		WhiteImage           string   `json:"white_image"`
		ZkFinalPrice         string   `json:"zk_final_price"`
	} `json:"result_list"`
	TotalResults int    `json:"total_results"`
	RequestID    string `json:"request_id"`
}

type TaobaoTbkItemConvertResponse struct {
	Results []struct {
		ClickURL string `json:"click_url"`
		NumIid   int64  `json:"num_iid"`
	} `json:"results"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkShopConvertResponse struct {
	Results []struct {
		ClickURL string `json:"click_url"`
		UserID   int    `json:"user_id"`
	} `json:"results"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkTpwdCreateResponse struct {
	Data struct {
		Model          string `json:"model"`
		PasswordSimple string `json:"password_simple"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkDgPunishOrderGetResponse struct {
	Result struct {
		BizErrorCode int `json:"biz_error_code"`
		Data         struct {
			PageNo   int `json:"page_no"`
			PageSize int `json:"page_size"`
			Results  []struct {
				PunishStatus      string `json:"punish_status"`
				RelationID        int64  `json:"relation_id"`
				SettleMonth       int    `json:"settle_month"`
				TbTradeID         int64  `json:"tb_trade_id"`
				TkAdzoneID        int64  `json:"tk_adzone_id"`
				TkPubID           string `json:"tk_pub_id"`
				TkSiteID          int    `json:"tk_site_id"`
				TkTradeCreateTime string `json:"tk_trade_create_time"`
				ViolationType     string `json:"violation_type"`
			} `json:"results"`
			TotalCount int `json:"total_count"`
		} `json:"data"`
		ResultCode int `json:"result_code"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkScPublisherInfoSaveResponse struct {
	Data struct {
		AccountName string `json:"account_name"`
		Desc        string `json:"desc"`
		RelationID  int64  `json:"relation_id"`
		SpecialID   int64  `json:"special_id"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkScPublisherInfoGetResponse struct {
	Data struct {
		InviterList []struct {
			RelationApp  string `json:"relation_app"`
			CreateDate   string `json:"create_date"`
			AccountName  string `json:"account_name"`
			RealName     string `json:"real_name"`
			RelationID   int    `json:"relation_id"`
			OfflineScene string `json:"offline_scene"`
			OnlineScene  string `json:"online_scene"`
			Note         string `json:"note"`
			RootPid      string `json:"root_pid"`
			Rtag         string `json:"rtag"`
			OfflineInfo  struct {
				ShopName        string `json:"shop_name"`
				ShopType        string `json:"shop_type"`
				PhoneNumber     string `json:"phone_number"`
				DetailAddress   string `json:"detail_address"`
				Location        string `json:"location"`
				ShopCertifyType string `json:"shop_certify_type"`
				CertifyNumber   string `json:"certify_number"`
				Career          string `json:"career"`
			} `json:"offline_info"`
			SpecialID    int    `json:"special_id"`
			PunishStatus string `json:"punish_status"`
			ExternalID   string `json:"external_id"`
		} `json:"inviter_list"`
		RootPidChannelList []string `json:"root_pid_channel_list"`
		TotalCount         int      `json:"total_count"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

type TaobaoTbkScInvitecodeGetResponse struct {
	Data struct {
		InviterCode string `json:"inviter_code"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}

type TaobaoTopAuthTokenCreateResponse struct {
	TokenResult TaobaoTopAuthTokenResult `json:"token_result"`
	RequestID   string                   `json:"request_id"`
}

type TaobaoTopAuthTokenRefreshResponse struct {
	TokenResult TaobaoTopAuthTokenResult `json:"token_result"`
	RequestID   string                   `json:"request_id"`
}

type TaobaoTopAuthResult struct {
	TokenResult string `json:"token_result"`
	RequestID   string `json:"request_id"`
}

type TaobaoTopAuthTokenResult struct {
	W1ExpiresIn           int    `json:"w1_expires_in"`
	RefreshTokenValidTime int64  `json:"refresh_token_valid_time"`
	TaobaoUserNick        string `json:"taobao_user_nick"`
	ReExpiresIn           int    `json:"re_expires_in"`
	ExpireTime            int64  `json:"expire_time"`
	TokenType             string `json:"token_type"`
	AccessToken           string `json:"access_token"`
	TaobaoOpenUID         string `json:"taobao_open_uid"`
	W1Valid               int64  `json:"w1_valid"`
	RefreshToken          string `json:"refresh_token"`
	W2ExpiresIn           int    `json:"w2_expires_in"`
	W2Valid               int64  `json:"w2_valid"`
	R1ExpiresIn           int    `json:"r1_expires_in"`
	R2ExpiresIn           int    `json:"r2_expires_in"`
	R2Valid               int64  `json:"r2_valid"`
	R1Valid               int64  `json:"r1_valid"`
	TaobaoUserID          string `json:"taobao_user_id"`
	ExpiresIn             int    `json:"expires_in"`
}

type TaobaoTbkItemClickExtractResponse struct {
	ItemID    string `json:"item_id"`
	OpenIid   string `json:"open_iid"`
	RequestID string `json:"request_id"`
}
