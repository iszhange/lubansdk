package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkRelationRefund(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkRelationRefundRequest()
	req.SetParam("search_type", "3")
	req.SetParam("refund_type", "0")
	req.SetParam("biz_type", "1")
	req.SetParam("start_time", "2020-11-09 00:00:00")
	req.SetParam("page_no", "1")
	req.SetParam("page_size", "3")
	body, err := c.Exec(req)
	if err != nil {
		t.Error(err)
	}

	result, err := req.Result(body)
	if err != nil {
		responseError, _ := top.ErrorResponse(body)
		t.Error(responseError)
	} else {
		fmt.Println(result)
	}
}
