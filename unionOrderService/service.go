package unionOrderService

import "github.com/varluffy/vipsdk"

type Service struct{}

func (a Service) ServiceName() string {
	return "com.vip.adp.api.open.service.UnionOrderService"
}

func (a Service) Version() string {
	return vipsdk.DefaultVersion
}

func (a Service) Token() bool {
	return false
}
type OrderInfo struct {
	OrderSn     string `json:"orderSn"` // 订单号
	Status      int    `json:"status"` // 订单状态:0-不合格，1-待定，2-已完结
	NewCustomer int    `json:"newCustomer"` // 新老客：0-待定，1-新客，2-老客
	ChannelTag  string `json:"channelTag"` // 渠道商模式下表示自定义渠道标识；工具商模式下表示pid
	OrderTime   int    `json:"orderTime"` // 下单时间 时间戳 单位毫秒
	SignTime    int    `json:"signTime"` // 签收时间 时间戳 单位毫秒
	SettledTime int    `json:"settledTime"` // 结算时间 时间戳 单位毫秒
	DetailList []*DetailList `json:"detailList"` // 商品明细
	LastUpdateTime            int    `json:"lastUpdateTime"` // 订单上次更新时间 时间戳 单位毫秒
	Settled                   int    `json:"settled"` // 订单结算状态 0-未结算,1-已结算
	SelfBuy                   int    `json:"selfBuy"` // 是否自推自买 0-否，1-是
	OrderSubStatusName        string `json:"orderSubStatusName"` //
	Commission                string `json:"commission"`
	AfterSaleChangeCommission string `json:"afterSaleChangeCommission"`
	AfterSaleChangeGoodsCount int    `json:"afterSaleChangeGoodsCount"`
	CommissionEnterTime       int    `json:"commissionEnterTime"`
	OrderSource               string `json:"orderSource"`
	Pid                       string `json:"pid"`
	IsPrepay                  int    `json:"isPrepay"`
	StatParam                 string `json:"statParam"`
}

// 商品明细
type DetailList struct {
	GoodsID             string `json:"goodsId"` // 商品id
	GoodsName           string `json:"goodsName"` // 商品名称
	GoodsThumb          string `json:"goodsThumb"` // 商品缩略图
	GoodsCount          int    `json:"goodsCount"` // 商品数量
	CommissionTotalCost string `json:"commissionTotalCost"` // 商品计佣金额(元,保留两位小数)
	CommissionRate      string `json:"commissionRate"` // 商品佣金比例(%)
	Commission          string `json:"commission"` //  商品佣金金额(元,保留两位小数)
	CommCode            string `json:"commCode"` // 佣金编码：对应商品二级分类
	CommName            string `json:"commName"` // 佣金方案名称
	OrderSource         string `json:"orderSource"` // 订单来源
	AfterSaleInfo []*AfterSaleInfo `json:"afterSaleInfo"` // 商品售后信息
	SizeID         string `json:"sizeId"` // 商品尺码：2019.01.01之后可用
	Status         int    `json:"status"` // 商品状态：0-不合格，1-待定，2-已完结
	BrandStoreSn   string `json:"brandStoreSn"` // 品牌编号
	BrandStoreName string `json:"brandStoreName"` // 品牌名称
}

// 商品售后信息
type AfterSaleInfo struct {
	AfterSaleChangedCommission string `json:"afterSaleChangedCommission"` // 商品佣金售后变动:仅在订单完结之后发生售后时返回，无售后时为空
	AfterSaleChangedGoodsCount int    `json:"afterSaleChangedGoodsCount"` // 商品数量售后变动:仅在订单完结之后发生售后时返回，无售后时为空
	AfterSaleSn                string `json:"afterSaleSn,omitempty"` // 商品售后单号，无售后时为空
	AfterSaleStatus            int    `json:"afterSaleStatus,omitempty"` // 商品售后状态：1-售后中，2-售后完成，3-售后取消，无售后时为空
	AfterSaleType              int    `json:"afterSaleType"` // 售后类型：1-退货，2-换货,无售后时为空
	AfterSaleFinishTime        int    `json:"afterSaleFinishTime"` // 售后完成时间，时间戳，单位：毫秒，无售后时为空
}