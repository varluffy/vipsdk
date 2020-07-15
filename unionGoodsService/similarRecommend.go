package unionGoodsService

type SimilarRecommendParam struct {
	Service
	Request SimilarRecommendRequest
}

func (a SimilarRecommendParam) MethodName() string {
	return "similarRecommend"
}

func (a SimilarRecommendParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["request"] = a.Request
	return p
}

// 相似推荐商品列表请求参数
type SimilarRecommendRequest struct {
	GoodsId                     string       `json:"goodsId"`                               // 相似推荐商品列表请求参数
	Page                        int64        `json:"page"`                                  // 页码
	PageSize                    int64        `json:"pageSize,omitempty"`                    // 分页大小:默认20，最大100
	RequestId                   string       `json:"requestId,omitempty"`                   // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	CommonParams                CommonParams `json:"commonParams,omitempty"`                // 通用参数
	QueryReputation             bool         `json:"queryReputation,omitempty"`             // 是否查询商品评价信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
	QueryStoreServiceCapability bool         `json:"queryStoreServiceCapability,omitempty"` // 是否查询店铺服务能力信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
	QueryStock                  bool         `json:"queryStock,omitempty"`                  // 是否查询库存:默认不查询
	QueryActivity               bool         `json:"queryActivity,omitempty"`               // 是否查询商品活动信息:默认不查询
	QueryPrepay                 bool         `json:"queryPrepay,omitempty"`                 // 是否查询商品预付信息:默认不查询
	ChanTag                     string       `json:"chanTag,omitempty"`                     // pid
}

type SimilarRecommendResponse struct {
	ReturnCode string            `json:"returnCode"`
	Result     GoodsInfoResponse `json:"result"`
}
