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
