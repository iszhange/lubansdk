package requests_test

import (
	"fmt"
	"lubantop"
	"testing"
)

func TestTaobaoTbkShopGet(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkShopGetRequest()
	req.SetFields("user_id,shop_title,shop_type,seller_nick,pict_url,shop_url")
	req.SetQ("手机")
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
