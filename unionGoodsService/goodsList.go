package unionGoodsService

type GoodsListParam struct {
	Service
	Request GoodsInfoRequest
	RequestId string
}

func (a GoodsListParam) MethodName() string {
	return "goodsList"
}

func (a GoodsListParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["request"] = a.Request
	return p
}

type GoodsInfoRequest struct {
	ChannelType int64 `json:"channelType,omitempty"` // 频道类型:0-超高佣，1-出单爆款; 当请求类型为频道时必传
	Page int64 `json:"page,omitempty"` // 页码
	PageSize int64 `json:"pageSize,omitempty"` // 分页大小:默认20，最大100
	RequestId string `json:"requestId,omitempty"` // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	QueryReputation bool `json:"queryReputation,omitempty"` // 是否查询商品评价信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
	QueryStoreServiceCapability bool `json:"queryStoreServiceCapability,omitempty"` // 是否查询店铺服务能力信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
	SourceType int64 `json:"sourceType,omitempty"` // 请求源类型：0-频道，1-组货
	JxCode string `json:"jxCode,omitempty"` // 精选组货码：当请求源类型为组货时必传
	FieldName string `json:"fieldName,omitempty"` // 排序字段: COMMISSION-佣金，PRICE-价格,COMM_RATIO-佣金比例，DISCOUNT-折扣
	Order int64 `json:"order,omitempty"` // 排序顺序：0-正序，1-逆序，默认正序
	QueryStock bool `json:"queryStock,omitempty"` // 是否查询库存:默认不查询
	QueryActivity bool `json:"queryActivity,omitempty"` // 是否查询商品活动信息:默认不查询
	QueryPrepay bool `json:"queryPrepay,omitempty"` // 是否查询商品预付信息:默认不查询
	CommonParams CommonParams `json:"commonParams,omitempty"` // 通用参数
	VendorCode string `json:"vendorCode,omitempty"` // 工具商code
	ChanTag string `json:"chanTag,omitempty"` // pid
}

type GoodsListResponse struct {
	ReturnCode string       `json:"returnCode"`
	Result     GoodsInfoResponse `json:"result"`
}