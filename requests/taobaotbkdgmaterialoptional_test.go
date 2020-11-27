package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubantop"
)

func TestTaobaoTbkDgMaterialOptional(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkDgMaterialOptionalRequest()
	req.SetQ("上衣")
	req.SetMaterialId(17004)
	req.SetAdzoneId("107356100348")
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
