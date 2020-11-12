package requests_test

import (
	"fmt"
	"lubantop"
	"testing"
)

func TestTaobaoTopAuthTokenCreate(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTopAuthTokenCreateRequest()
	req.SetCode("7oBgi6L31Vr1Gk9wbT7T4Naf9071114")
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
