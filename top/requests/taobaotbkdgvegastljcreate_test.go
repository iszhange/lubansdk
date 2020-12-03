package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkDgVegasTljCreate(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkDgVegasTljCreateRequest()
	req.SetParam("adzone_id", "107356100348")
	req.SetParam("item_id", "621069848303")
	req.SetParam("total_num", "1")
	req.SetParam("name", "测试")
	req.SetParam("user_total_win_num_limit", "1")
	req.SetParam("security_switch", "true")
	req.SetParam("per_face", "1")
	req.SetParam("send_start_time", "2020-11-11 00:00:00")
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
