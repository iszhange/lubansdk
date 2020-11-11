package requests_test

import (
	"fmt"
	"lubantop"
	"testing"
)

func TestTaobaoTbkShopConvert(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkShopConvertRequest()
	req.SetFields("user_id,click_url")
	req.SetUserIds("748612647")
	req.SetAdzoneId("107356100348")
	req.SetUnid("xxx")
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
