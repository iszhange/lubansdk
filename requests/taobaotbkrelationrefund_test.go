package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubantop"
)

func TestTaobaoTbkRelationRefund(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkRelationRefundRequest()
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
		responseError, _ := lubantop.ErrorResponse(body)
		t.Error(responseError)
	} else {
		fmt.Println(result)
	}
}
