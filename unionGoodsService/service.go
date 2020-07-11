package unionGoodsService

import "github.com/varluffy/vipsdk"

type Service struct{
	RequestId string `json:"requestId"` // 请求id：UUID
}

func (a Service) ServiceName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

func (a Service) Version() string {
	return vipsdk.DefaultVersion
}

func (a Service) Token() bool {
	return false
}

type GoodsInfo struct {
	GoodsId                string   `json:"goodsId,omitempty"`
	GoodsName              string   `json:"goodsName,omitempty"`
	GoodsDesc              string   `json:"goodsDesc,omitempty"`
	DestURL                string   `json:"destUrl,omitempty"`
	GoodsThumbURL          string   `json:"goodsThumbUrl,omitempty"`
	GoodsCarouselPictures  []string `json:"goodsCarouselPictures,omitempty"`
	GoodsMainPicture       string   `json:"goodsMainPicture,omitempty"`
	CategoryId             int64    `json:"categoryId,omitempty"`
	CategoryName           string   `json:"categoryName,omitempty"`
	SourceType             int64    `json:"sourceType,omitempty"`
	MarketPrice            string   `json:"marketPrice,omitempty"`
	VipPrice               string   `json:"vipPrice,omitempty"`
	CommissionRate         string   `json:"commissionRate,omitempty"`
	Commission             string   `json:"commission,omitempty"`
	Discount               string   `json:"discount,omitempty"`
	GoodsDetailPictures    []string `json:"goodsDetailPictures,omitempty"`
	Cat1stId               int64    `json:"cat1stId,omitempty"`
	Cat1stName             string   `json:"cat1stName,omitempty"`
	Cat2ndId               int64    `json:"cat2ndId,omitempty"`
	Cat2ndName             string   `json:"cat2ndName,omitempty"`
	BrandStoreSn           string   `json:"brandStoreSn,omitempty"`
	BrandName              string   `json:"brandName,omitempty"`
	BrandLogoFull          string   `json:"brandLogoFull,omitempty"`
	SchemeEndTime          int64    `json:"schemeEndTime,omitempty"`
	SellTimeFrom           int64    `json:"sellTimeFrom,omitempty"`
	SellTimeTo             int64    `json:"sellTimeTo,omitempty"`
	Weight                 int64    `json:"weight,omitempty"`
	StoreInfo              StoreInfo
	CommentsInfo           GoodsCommentsInfo
	storeServiceCapability StoreServiceCapability
	BrandId                int64 `json:"brandId,omitempty"`
	SchemeStartTime        int64 `json:"schemeStartTime,omitempty"`
	SaleStockStatus        int64 `json:"saleStockStatus,omitempty"`
	Status                 int64 `json:"status,omitempty"`
	PrepayInfo             PrepayInfo
	JoinedActivities       []*ActivityInfo
	CouponInfo             PMSCouponInfo
}

// 店铺信息
type StoreInfo struct {
	StoreId   string `json:"storeId,omitempty"`   // 店铺id
	StoreName string `json:"storeName,omitempty"` // 店铺名称
}

// 商品评价信息
type GoodsCommentsInfo struct {
	Comments          int64  `json:"comments,omitempty"`          // 商品评论数
	GoodCommentsShare string `json:"goodCommentsShare,omitempty"` // 商品好评率:百分比，不带百分号
}

// 商品所属店铺服务能力评价
type StoreServiceCapability struct {
	StoreScore    string `json:"storeScore,omitempty"`    // 店铺评分：保留两位小数
	StoreRankRate string `json:"storeRankRate,omitempty"` // 店铺同品类排名比例：例如10表示前10%
}

// 商品预付信息
type PrepayInfo struct {
	IsPrepay             int64  `json:"isPrepay,omitempty"`             // 是否预付商品:0-否，1-是
	PrepayPrice          string `json:"prepayPrice,omitempty"`          // 预付到手价：元
	FirstAmount          string `json:"firstAmount,omitempty"`          // 预付首款金额：元
	LastAmount           string `json:"lastAmount,omitempty"`           // 预付尾款金额：元
	PrepayFavAmount      string `json:"prepayFavAmount,omitempty"`      // 预付优惠金额: 元
	DeductionPrice       string `json:"deductionPrice,omitempty"`       // 抵扣价格(首款+优惠金额)：元
	PrepayDiscount       string `json:"prepayDiscount,omitempty"`       // 预付折扣：(唯品价-优惠金额)/唯品价 保留两位小数的数字字符串
	PrepayFirstStartTime int64  `json:"prepayFirstStartTime,omitempty"` // 首款支付开始时间:时间戳,单位毫秒
	PrepayFirstEndTime   int64  `json:"prepayFirstEndTime,omitempty"`   // 首款支付结束时间:时间戳,单位毫秒
	PrepayLastStartTime  int64  `json:"prepayLastStartTime,omitempty"`  // 尾款支付开始时间:时间戳,单位毫秒
	PrepayLastEndTime    int64  `json:"prepayLastEndTime,omitempty"`              // 尾款支付结束时间:时间戳,单位毫秒
}

// 商品参加的唯品会活动信息
type ActivityInfo struct {
	ActType           int64  `json:"actType,omitempty"`           // 活动类型:18-唯品快抢
	ActName           string `json:"actName,omitempty"`           // 活动名称
	BeginTime         int64  `json:"beginTime,omitempty"`         // 开始时间：时间戳，单位毫秒
	EndTime           int64  `json:"endTime,omitempty"`           // 结束时间：时间戳，单位毫秒
	ForeShowBeginTime int64  `json:"foreShowBeginTime,omitempty"` // 预热开始时间：时间戳，单位毫秒
}
// 红包信息
type PMSCouponInfo struct {
	CouponNo   string `json:"couponNo,omitempty"` // 优惠券批次号
	CouponName string `json:"couponName,omitempty"` // 优惠劵名称
	Buy        string `json:"buy,omitempty"` // 使用门槛
	Fav        string `json:"fav,omitempty"` // 优惠金额
}

type SortFields struct {
	FieldName string `json:"fieldName,omitempty"`
	FieldDesc string `json:"fieldDesc,omitempty"`
}
// 通用参数
type CommonParams struct {
	Plat int64 `json:"plat,omitempty"` // 用户平台：1-PC,2-APP,3-小程序,不传默认为APP
	DeviceType string `json:"deviceType,omitempty"` // 设备号类型：IMEI，IDFA，OAID，有则传入
	DeviceValue string `json:"deviceValue,omitempty"` // 设备号MD5加密后的值，有则传入
	IP string `json:"ip,omitempty"` // 用户ip地址
	Longitude string `json:"longitude,omitempty"` // 经度 如:29.590961456298828
	Latitude string `json:"latitude,omitempty"` // 纬度 如:106.51573181152344
}

type GoodsInfoResponse struct {
	GoodsInfoList []*GoodsInfo `json:"goodsInfoList"`
	Total int64 `json:"total"`
	SortFields []*SortFields `json:"sortFields"`
	Page int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}
