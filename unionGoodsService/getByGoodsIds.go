package unionGoodsService

import (
	"net/url"
)

type GetByGoodsIdsParam struct {
	Service
	Request GetByGoodsIdsRequest
}

func (a GetByGoodsIdsParam) MethodName() string {
	return "getByGoodsIds"
}

func (a GetByGoodsIdsParam) Params() interface{} {
	p := make(url.Values)
	p["goodsIdList"] = a.Request.GoodsIdList
	p.Add("requestId", a.Request.RequestId)
	if a.Request.ChanTag != "" {
		p.Add("chanTag", a.Request.ChanTag)
	}
	return p
}

type GetByGoodsIdsRequest struct {
	GoodsIdList []string `json:"goodsIdList"` // 商品id列表
	ChanTag     string   `json:"chanTag"`     // 自定义渠道标识,同推广位 非必填
	RequestId   string   `json:"requestId"`
}
type GetByGoodsIdsResponse struct {
	ReturnCode string       `json:"returnCode"`
	Result     []*GoodsInfo `json:"result"`
}
