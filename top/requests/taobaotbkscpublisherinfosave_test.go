package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkScPublisherInfoSave(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret
	c.Session = "6202706416584egi88b06e33e463787f04b30dbd2ce59872418252546"

	req := top.TaobaoTbkScPublisherInfoSaveRequest()
	req.SetInviterCode("8W46PX")
	req.SetInfoType(1)
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
