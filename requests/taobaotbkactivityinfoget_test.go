package requests_test

import (
	"fmt"
	"lubantop"
	"testing"
)

func TestTaobaoTbkActivityInfoGet(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkActivityInfoGetRequest()
	req.SetAdzoneId(107356100348)
	req.SetActivityMaterialId("20150318020002040")
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
