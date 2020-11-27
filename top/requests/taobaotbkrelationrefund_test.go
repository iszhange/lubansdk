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
	req.SetSearchType(3)
	req.SetRefundType(0)
	req.SetBizType(1)
	req.SetStartTime("2020-11-09 00:00:00")
	req.SetPageNo(1)
	req.SetPageSize(3)
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
