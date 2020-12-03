package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkScVegasSendReport(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret
	c.Session = "6202a07eb63bfbbdf74dcdd7a74219930d4046114b8718a527712929"

	req := top.TaobaoTbkScVegasSendReportRequest()
	req.SetParam("biz_date", "20201101")
	req.SetParam("activity_id", "1306")
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
