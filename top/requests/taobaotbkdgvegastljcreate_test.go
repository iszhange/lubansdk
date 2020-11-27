package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkDgVegasTljCreate(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkDgVegasTljCreateRequest()
	req.SetAdzoneId("107356100348")
	req.SetItemId("621069848303")
	req.SetTotalNum(1)
	req.SetName("测试")
	req.SetUserTotalWinNumLimit(1)
	req.SetSecuritySwitch(true)
	req.SetPerFace("1")
	req.SetSendStartTime("2020-11-11 00:00:00")
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
