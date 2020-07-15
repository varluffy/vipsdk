package UnionUrlService

type VipLinkCheckParam struct {
	Service
	Request VipLinkCheckReq
}

func (a VipLinkCheckParam) MethodName() string  {
	return "vipLinkCheck"
}

func (a VipLinkCheckParam) Params() interface{}  {
	p := make(map[string]interface{}, 0)
	p["vipLinkCheckReq"] = a.Request
	return p
}

// 链接解析请求
type VipLinkCheckReq struct {
	Source string`json:"source,omitempty"` // 请求源：调用方自定义，标识请求应用即可，无业务含义
	Content string `json:"content,omitempty"` // 检查的内容(长度不超过10000)
}


type VipLinkCheckResponse struct {
	ReturnCode string `json:"returnCode"`
	Result struct {
		SuccessMap map[string]VipLinkCheckVO `json:"successMap"`
		FailMap map[string]VipLinkCheckVO `json:"failMap"`
	} `json:"result"`
}

// 解析map，key为原始url
type VipLinkCheckVO struct {
	LinkType string `json:"linkType"`
	LandURL  string `json:"landUrl"`
	GoodsID  string `json:"goodsId"`
	BrandID  string `json:"brandId"`
}
