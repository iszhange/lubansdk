package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkScInvitecodeGet(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret
	c.Session = "6200326740fa017c194a82bc0b2fc7fee2ZZ5754c3af44d1076725086"

	req := top.TaobaoTbkScInvitecodeGetRequest()
	req.SetParam("relation_app", "common")
	req.SetParam("code_type", "1")
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
