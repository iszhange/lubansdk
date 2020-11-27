package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkActivityInfoGet(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkActivityInfoGetRequest()
	req.SetAdzoneId(107356100348)
	req.SetActivityMaterialId("20150318020002040")
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
