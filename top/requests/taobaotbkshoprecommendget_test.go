package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkShopRecommendGet(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkShopRecommendGetRequest()
	req.SetFields("user_id,shop_title,shop_type,seller_nick,pict_url,shop_url")
	req.SetUserId(92688455)
	req.SetCount(4)
	req.SetPlatform(1)
	body, err := c.Exec(req)
	if err != nil {
		t.Error(err)
	}

	result, err := req.Result(body)
	if err != nil {
		responseError, _ := top.ErrorResponse(body)
		t.Error(responseError.ErrorResponse.Code, responseError.ErrorResponse.Msg)
	} else {
		fmt.Println(result)
	}
}
