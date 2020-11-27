package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubantop"
)

func TestTaobaoTbkScInvitecodeGet(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret
	c.Session = "6200326740fa017c194a82bc0b2fc7fee2ZZ5754c3af44d1076725086"

	req := lubantop.TaobaoTbkScInvitecodeGetRequest()
	req.SetRelationApp("common")
	req.SetCodeType(1)
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
