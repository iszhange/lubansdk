package requests_test

import (
	"fmt"
	"testing"

	"github.com/iszhange/lubansdk/top"
)

func TestTaobaoJuItemsSearch(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoJuItemsSearchRequest()
	req.SetPid("mm_109757283_218800301_107356100348")
	req.SetCurrentPage(1)
	req.SetPageSize(20)
	req.SetWord("显示器")
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
