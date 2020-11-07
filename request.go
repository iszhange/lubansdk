package lubantop

import "lubantop/requests"

func TaobaoTbkItemInfoGetRequest() *requests.TaobaoTbkItemInfoGet {
	var r requests.TaobaoTbkItemInfoGet
	return r.New()
}
