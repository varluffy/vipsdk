package unionUrlService

type GenByVIPUrlParam struct {
	Service
	Request GenByVIPUrlRequest
}

func (a GenByVIPUrlParam) MethodName() string {
	return "genByVIPUrl"
}

func (a GenByVIPUrlParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["urlList"] = a.Request.UrlList

	p["requestId"] = a.Request.RequestId
	if a.Request.ChanTag != "" {
		p["chanTag"] = a.Request.ChanTag
	}
	if a.Request.StatParam != "" {
		p["statParam"] = a.Request.StatParam
	}
	return p
}

type GenByVIPUrlRequest struct {
	UrlList   []string `json:"urlList,omitempty"` // 商品id列表
	ChanTag   string   `json:"chanTag"`           // 渠道标识
	RequestId string   `json:"requestId"`         // 请求id：UUID
	StatParam string   `json:"statParam"`         // 自定义渠道统计参数
}

type GenByVIPUrlResponse struct {
	ReturnCode string `json:"returnCode"`
	Result     struct {
		URLInfoList []*UrlInfo `json:"urlInfoList"`
	} `json:"result"`
}
