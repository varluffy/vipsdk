package unionUrlService

type GenByGoodsIdParam struct {
	Service
	Request GenByGoodsIdRequest
}

func (a GenByGoodsIdParam) MethodName() string {
	return "genByGoodsId"
}

func (a GenByGoodsIdParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["goodsIdList"] = a.Request.GoodsIdList
	p["requestId"] = a.Request.RequestId
	if a.Request.ChanTag != "" {
		p["chanTag"] = a.Request.ChanTag
	}
	if a.Request.StatParam != "" {
		p["statParam"] = a.Request.StatParam
	}
	return p
}

type GenByGoodsIdRequest struct {
	GoodsIdList []string `json:"goodsIdList,omitempty"` // 商品id列表
	ChanTag     string   `json:"chanTag"`               // 渠道标识
	RequestId   string   `json:"requestId"`             // 请求id：UUID
	StatParam   string   `json:"statParam"`             // 自定义渠道统计参数
}

type UrlGenResponse struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	Result        struct {
		URLInfoList []*UrlInfo `json:"urlInfoList"`
	} `json:"result"`
}

type UrlInfo struct {
	Source         string `json:"source"`
	URL            string `json:"url"`
	LongURL        string `json:"longUrl"`
	UlURL          string `json:"ulUrl"`
	DeepLinkURL    string `json:"deeplinkUrl"`
	TraFrom        string `json:"traFrom"`
	NoEvokeURL     string `json:"noEvokeUrl"`
	NoEvokeLongURL string `json:"noEvokeLongUrl"`
	VipWxURL       string `json:"vipWxUrl"`
	VipWxCode      string `json:"vipWxCode"`
	VipQuickAppURL string `json:"vipQuickAppUrl"`
}
