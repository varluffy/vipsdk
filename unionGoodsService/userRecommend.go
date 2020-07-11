package unionGoodsService

type UserRecommendParam struct {
	Service
	Request UserRecommendRequest
}

func (a UserRecommendParam) MethodName() string {
	return "userRecommend"
}

func (a UserRecommendParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["request"] = a.Request
	return p
}

type UserRecommendRequest struct {
	Page int64 `json:"page,omitempty"` // 页码
	PageSize int64 `json:"pageSize,omitempty"` // 分页大小:默认20，最大100
	RequestId string `json:"requestId,omitempty"` // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	InStock int64 `json:"inStock"` // 是否有货 1:有货 0:无货 默认1
	GoodsSaleStats int64 `json:"goodsSaleStats"` // 商品售卖状态 1（在售）， 2（预热）， 3（在售+预热） 默认1
	OfflineStore int64 `json:"offlineStore"` // 筛选线下店商品：1只返线下店，0或者不传只返回特卖会
	CommonParams CommonParams `json:"commonParams,omitempty"` // 通用参数
	ChanTag string `json:"chanTag"` // 自定义渠道标识
}

type UserRecommendResponse struct {
	ReturnCode string       `json:"returnCode"`
	Result     GoodsInfoResponse `json:"result"`
}
