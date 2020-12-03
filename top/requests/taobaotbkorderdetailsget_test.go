package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkOrderDetailsGet(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkOrderDetailsGetRequest()
	req.SetParam("start_time", "2020-11-09 00:00:00")
	req.SetParam("end_time", "2020-11-09 00:10:00")
	req.SetParam("order_scene", "2")
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
