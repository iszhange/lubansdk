package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubantop"
)

func TestTaobaoTbkCouponGet(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkCouponGetRequest()
	req.SetItemId("619028455482")
	req.SetActivityId("1f7f93f7473f45bcab47aede5b12af3f")
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
