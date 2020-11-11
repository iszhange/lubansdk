package requests_test

import (
	"fmt"
	"lubantop"
	"testing"
)

func TestTaobaoTbkShopRecommendGet(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkShopRecommendGetRequest()
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
		responseError, _ := lubantop.ErrorResponse(body)
		t.Error(responseError.ErrorResponse.Code, responseError.ErrorResponse.Msg)
	} else {
		fmt.Println(result)
	}
}
