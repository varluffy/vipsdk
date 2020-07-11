package unionOrderService

type RefundOrderListParam struct {
	Service
	Request RefundOrderRequest
}

func (a RefundOrderListParam) MethodName() string {
	return "refundOrderList"
}

func (a RefundOrderListParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["request"] = a.Request
	return p
}

// 联盟维权订单请求参数
type RefundOrderRequest struct {
	SearchType int64 `json:"searchType"` // 查询类型:0-维权发起时间，1-维权完成时间（佣金扣除入账时间），2-佣金扣除结算时间
	SearchTimeStart int64 `json:"searchTimeStart"` // 目标时间起始：时间戳，单位毫秒; 当searchType=0时，为维权发起时间， 当searchType=1时，为维权完成时间（佣金扣除入账时间）， 当searchType=2时，为佣金扣除结算时间
	SearchTimeEnd int64 `json:"searchTimeEnd"` // 目标时间结束：时间戳，单位毫秒; 当searchType=0时，为维权发起时间， 当searchType=1时，为维权完成时间（佣金扣除入账时间）， 当searchType=2时，为佣金扣除结算时间
	OrderSns []string `json:"orderSns"` // 目标订单号集合:当指定订单号集合时，目标时间区间可以不传,否则必须指定目标时间区间
	Page int64 `json:"page,omitempty"` // 页码
	PageSize int64 `json:"pageSize,omitempty"` // 分页大小:默认20，最大100
	RequestId string `json:"requestId,omitempty"` // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
}
//
type RefundOrderResponse struct {
	ReturnCode string `json:"returnCode"`
	Result struct{
		RefundOrderInfoList []*RefundOrderInfo `json:"refundOrderInfoList"` // 联盟维权订单信息
		Total int64 `json:"total"` // 总记录数
		Page int64 `json:"page"` // 当前页
		PageSize int64 `json:"pageSize"` // 分页大小
	} `json:"result"`
}

// 联盟维权订单信息
type RefundOrderInfo struct {
	OrderSn               string `json:"orderSn"` // 订单号
	ApplyTime             int    `json:"applyTime"` // 申请时间:时间戳，单位毫秒
	RefundType            int    `json:"refundType"` // 维权类型：1-退货，2-换货
	CommissionTotalCost   string `json:"commissionTotalCost"` // 订单计佣金额(元,两位小数)
	Commission            string `json:"commission"` // 订单佣金(元,两位小数)
	GoodsCount            int    `json:"goodsCount"` // 商品数量
	CommissionEnterTxn    string `json:"commissionEnterTxn"` // 维权返还佣金的入账流水号
	CommissionEnterTime   int    `json:"commissionEnterTime"` // 维权返还佣金的入账时间: 时间戳，单位毫秒
	CommissionSettledTime int    `json:"commissionSettledTime"` // 维权返回佣金的结算时间：时间戳，单位毫秒
	RefundOrderDetails    []*RefundOrderDetail `json:"refundOrderDetails"` // 维权订单明细
	UserID              int    `json:"userId"` // 订单归属人id
	OrderTime           int    `json:"orderTime"` // 订单时间:时间戳，单位毫秒
	OriginCommEnterTime int    `json:"originCommEnterTime"` // 订单入账时间：发起维权之前入账时间，时间戳，单位毫秒
	OriginCommEnterTxn  string `json:"originCommEnterTxn"` // 订单入账流水：发起维权之前入账流水号
	Settled             int    `json:"settled"` // 售后订单结算状态：0-未结算，1-已结算
	AfterSaleSn         string `json:"afterSaleSn"` // 售后单号
	AfterSaleStatus     int    `json:"afterSaleStatus"` // 售后单状态:1-售后中，2-售后完成，3-售后取消
}

// 维权订单明细
type RefundOrderDetail struct {
	GoodsID    string `json:"goodsId"` // 商品id
	SizeID     string `json:"sizeId"` // 商品尺码id
	GoodsPrice string `json:"goodsPrice"` // 商品价格：元
	GoodsCount int    `json:"goodsCount"` // 维权商品数量
	Commission string `json:"commission"` // 维权商品佣金
	TotalCost  string `json:"totalCost"` // 维权商品计佣金额
}