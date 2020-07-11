package unionOrderService

type OrderListParam struct {
	Service
	Request OrderQueryModel
}

func (a OrderListParam) MethodName() string {
	return "orderList"
}

func (a OrderListParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["queryModel"] = a.Request
	return p
}

type OrderQueryModel struct {
	Status int64 `json:"status,omitempty"` // 订单状态:0-不合格，1-待定，2-已完结，该参数不设置默认代表全部状态
	OrderTimeStart int64 `json:"orderTimeStart,omitempty"` // 订单时间起始 时间戳 单位毫秒
	OrderTimeEnd int64 `json:"orderTimeEnd,omitempty"` // 订单时间结束 时间戳 单位毫秒
	Page int64 `json:"page,omitempty,omitempty"` // 页码
	PageSize int64 `json:"pageSize,omitempty"` // 分页大小:默认20，最大100
	RequestId string `json:"requestId,omitempty"` // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	UpdateTimeStart int64 `json:"updateTimeStart,omitempty"` // 更新时间-起始 时间戳 单位毫秒
	UpdateTimeEnd int64 `json:"updateTimeEnd,omitempty"` // 下单时间-结束 时间戳 单位毫秒
	OrderSnList []string `json:"orderSnList,omitempty"` // 订单号列表：当传入订单号列表时，订单时间和更新时间区间可不传入
	VendorCode string `json:"vendorCode,omitempty"` // 工具商code
	ChanTag string `json:"chanTag,omitempty"` // pid
}

type OrderListResponse struct {
	ReturnCode string `json:"returnCode"`
	Result struct{
		OrderInfoList []*OrderInfo `json:"orderInfoList"`
		Total int64 `json:"total"`
		Page int64 `json:"page"`
		PageSize int64 `json:"pageSize"`
	} `json:"result"`

}