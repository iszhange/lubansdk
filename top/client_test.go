package top_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestClient(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	for i := 0; i < 150; i++ {
		req := top.TaobaoTbkOrderDetailsGetRequest()
		req.SetParam("start_time", "2020-05-27 11:00:00")
		req.SetParam("end_time", "2020-05-27 11:10:00")
		req.SetParam("order_scene", "2")
		body, err := c.Exec(req)
		if err != nil {
			log.Fatalln(err)
		}
		result, err := req.Result(body)
		if err != nil {
			responseError, _ := top.ErrorResponse(body)
			t.Error(responseError)
		} else {
			fmt.Println(result)
		}
	}

}
