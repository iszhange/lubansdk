package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoTbkSpreadGet(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkSpreadGetRequest()
	req.SetParam("url", "https://uland.taobao.com/coupon/edetail?spm=a2e1u.13363363.20282407.11.635a7466h7jfhf&e=YPXgq7onIj0NfLV8niU3RxrSI%2FOabn6qNg4Gqf8CT4BnmB%2Fzds2ljTK%2Fvbppw1RAkTBHbvZZmkcAbAN3NcqP1djiLvKBoQ2vzOJz0K1X%2B2oyv726acNUQA7XsmpjBykO26z7jT1gC8I5ojmj6WiwtslMTjSQsK4p5bgGAUYJbG0%3D&app_pvid=59590_33.5.224.5_599_1605017577531&ptl=floorId%3A23919%3Bapp_pvid%3A59590_33.5.224.5_599_1605017577531%3Btpp_pvid%3A1c7a6ae3-96bb-47fe-b6ee-66d9246e8e42&union_lens=lensId%3AOPT%401605017577%401c7a6ae3-96bb-47fe-b6ee-66d9246e8e42_529031840597%401%3Brecoveryid%3A201_11.88.26.84_3329608_1605017577031%3Bprepvid%3A201_11.88.26.84_3329608_1605017577031&pid=mm_10011550_0_0")
	req.SetParam("url", "https://uland.taobao.com/coupon/edetail?spm=a2e1u.13363363.20282407.19.635a7466h7jfhf&e=ZWq90cyT6CUNfLV8niU3RxrSI%2FOabn6qNg4Gqf8CT4BnmB%2Fzds2ljS9%2Fp8yJYuNZ%2Bac4e%2FdmuMcAbAN3NcqP1djiLvKBoQ2vzOJz0K1X%2B2ovf6fMiWLjWRIM4feFv3RlgxcDqqH5BbzG4hYZmNyX%2BQrE199RXvLx%2FjtLOxCEadVVbrKqp4Yn8g%3D%3D&app_pvid=59590_33.5.224.5_599_1605017577531&ptl=floorId%3A23919%3Bapp_pvid%3A59590_33.5.224.5_599_1605017577531%3Btpp_pvid%3A1c7a6ae3-96bb-47fe-b6ee-66d9246e8e42&union_lens=lensId%3AOPT%401605017577%401c7a6ae3-96bb-47fe-b6ee-66d9246e8e42_602863139190%401%3Brecoveryid%3A201_11.88.26.84_3329608_1605017577031%3Bprepvid%3A201_11.88.26.84_3329608_1605017577031&pid=mm_10011550_0_0")
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
