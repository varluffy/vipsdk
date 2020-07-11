package unionGoodsService

type QueryGoodsParam struct {
	Service
	Request QueryGoodsRequest
}

func (a QueryGoodsParam) MethodName() string {
	return "query"
}

func (a QueryGoodsParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["request"] = a.Request
	return p
}

// 关键词搜索商品请求参数
type QueryGoodsRequest struct {
	Keyword string `json:"keyword"` // 关键词
	FieldName string `json:"fieldName,omitempty"` // 排序字段
	Order int64 `json:"order,omitempty"` // 排序顺序：0-正序，1-逆序，默认正序
	Page int64 `json:"page,omitempty"` // 页码
	PageSize int64 `json:"pageSize,omitempty"` // 分页大小:默认20，最大100
	RequestId string `json:"requestId,omitempty"` // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	PriceStart string `json:"priceStart"` // 价格区间---start
	PriceEnd string `json:"priceEnd"` // 价格区间---end
	QueryReputation bool `json:"queryReputation,omitempty"` // 是否查询商品评价信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
	QueryStoreServiceCapability bool `json:"queryStoreServiceCapability,omitempty"` // 是否查询店铺服务能力信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
	QueryStock bool `json:"queryStock,omitempty"` // 是否查询库存:默认不查询
	QueryActivity bool `json:"queryActivity,omitempty"` // 是否查询商品活动信息:默认不查询
	QueryPrepay bool `json:"queryPrepay,omitempty"` // 是否查询商品预付信息:默认不查询
	CommonParams CommonParams `json:"commonParams,omitempty"` // 通用参数
	VendorCode string `json:"vendorCode,omitempty"` // 工具商code
	ChanTag string `json:"chanTag,omitempty"` // pid
}

type QueryGoodsResponse struct {
	ReturnCode string       `json:"returnCode"`
	Result     GoodsInfoResponse `json:"result"`
}


