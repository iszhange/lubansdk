package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkTpwdCreate(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkTpwdCreateRequest()
	req.SetParam("text", "测试生成Tpwd")
	req.SetParam("url", "https://uland.taobao.com/coupon/edetail?e=GQrLLm8sFLIGQASttHIRqWwrwhy1dV/8UojZmUyKjr5I/GMtm9ep/ce86G64Kn5VxmjDqrTQPvO5E1myvayfWUTiStE+0TIIMSEmlRStydiKyIANbW8yPApvmcNznetvdj3Hchy1wJYuynEbLfMGNjwh43yW+FEDaXOtHj5+txuO/1gtyYLZ+qOpNqvGxomLcW0FAk4P4CIcD1Tr03a5GA==&traceId=0bb61b5616050868757051422e6316&relationId=532835060&union_lens=lensId:TAPI@1605086875@0b57b76e_2670_175b6a21054_53a6@01&activityId=6a6eb1871c324374a84c26557056cb36&un=352487b97a007785b83707894294ed3d&share_crt_v=1&ut_sk=1.utdid_24796432_1605086877326.TaoPassword-Outside.taoketop&spm=a211b4.23434152&sp_tk=NnhWbGNQQUVnS0U=&bxsign=tcdQA-TA9sWTXsU5I05gic56pcijygXYf5-nrWc57hb8FKWxV_7Xx4kMqE3UjdKbwWgVamcAPdefKpx6C6GbiattF3FLpEb6-uEx-heiTNNWkI&visa=13a09278fde22a2e&disablePopup=true&disableSJ=1")
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
