package requests_test

import (
	"fmt"
	"lubantop"
	"testing"
)

func TestTaobaoTbkScVegasSendReport(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret
	c.Session = "6202a07eb63bfbbdf74dcdd7a74219930d4046114b8718a527712929"

	req := lubantop.TaobaoTbkScVegasSendReportRequest()
	req.SetBizDate("20201101")
	req.SetActivityId(1306)
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
