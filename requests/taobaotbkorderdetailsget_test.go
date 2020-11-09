package requests_test

import (
	"fmt"
	"lubantop"
	"testing"
)

func TestTaobaoTbkOrderDetailsGet(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	req := lubantop.TaobaoTbkOrderDetailsGetRequest()
	req.SetStartTime("2020-11-09 00:00:00")
	req.SetEndTime("2020-11-09 00:10:00")
	req.SetOrderScene(2)
	body, err := c.Exec(req)
	if err != nil {
		t.Error(err)
	}

	result, err := req.Result(body)
	if err != nil {
		responseError, _ := lubantop.ErrorResponse(body)
		t.Error(responseError)
	}

	fmt.Println(result)
}
