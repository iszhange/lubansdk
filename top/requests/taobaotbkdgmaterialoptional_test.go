package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkDgMaterialOptional(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkDgMaterialOptionalRequest()
	req.SetQ("上衣")
	req.SetMaterialId(17004)
	req.SetAdzoneId("107356100348")
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
