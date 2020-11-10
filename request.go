package lubantop

import "lubantop/requests"

func TaobaoTbkItemInfoGetRequest() *requests.TaobaoTbkItemInfoGet {
	var r requests.TaobaoTbkItemInfoGet
	return r.New()
}

func TaobaoTbkOrderDetailsGetRequest() *requests.TaobaoTbkOrderDetailsGet {
	var r requests.TaobaoTbkOrderDetailsGet
	return r.New()
}

func TaobaoTbkRelationRefundRequest() *requests.TaobaoTbkRelationRefund {
	var r requests.TaobaoTbkRelationRefund
	return r.New()
}

func TaobaoTbkDgVegasSendStatusRequest() *requests.TaobaoTbkDgVegasSendStatus {
	var r requests.TaobaoTbkDgVegasSendStatus
	return r.New()
}

func TaobaoTbkDgVegasSendReportRequest() *requests.TaobaoTbkDgVegasSendReport {
	var r requests.TaobaoTbkDgVegasSendReport
	return r.New()
}

func TaobaoTbkScVegasSendReportRequest() *requests.TaobaoTbkScVegasSendReport {
	var r requests.TaobaoTbkScVegasSendReport
	return r.New()
}
