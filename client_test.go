package lubantop_test

import (
	"log"
	"lubantop"
	"lubantop/requests"
	"testing"
)

func TestClient(t *testing.T) {
	c := lubantop.New()
	c.AppKey = lubantop.AppKey
	c.AppSecret = lubantop.AppSecret

	for i := 0; i < 150; i++ {
		req := requests.TbkOrderDetailsGetRequest()
		req.SetStartTime("2020-05-27 11:00:00")
		req.SetEndTime("2020-05-27 11:10:00")
		req.SetOrderScene(2)
		body, err := c.Exec(req)
		if err != nil {
			log.Fatalln(err)
		}
		var result requests.TbkOrderDetailsGetData
		result, err = requests.TbkOrderDetailsGetResult(body)
		if err != nil {
			if err.Error() == "接口返回错误" {
				errRet, err := requests.TbkOrderDetailGetError(body)
				if err == nil {
					log.Println(errRet)
				}
			} else {
				log.Fatalln(err)
			}
		}
		log.Println(result.Results)
	}

}
