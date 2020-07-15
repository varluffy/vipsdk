package unionPidService

type QueryPidParam struct {
	Service
	Request PidQueryRequest
}

func (a QueryPidParam) MethodName() string {
	return "queryPid"
}

func (a QueryPidParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["pidQueryRequest"] = a.Request
	return p
}

// 推广位查询请求信息
type PidQueryRequest struct {
	PidList   []string `json:"pidList,omitempty"`   // 推广为Id。 该参数不传时，会返回该userId下对应的所有的pid信息列表。
	RequestId string   `json:"requestId,omitempty"` // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	Page      int      `json:"page,omitempty"`      // 页码
	PageSize  int      `json:"pageSize,omitempty"`  // 页面大小：默认100
}

type CpsUnionPidQueryResponse struct {
	ReturnCode string `json:"returnCode"`
	Result     struct {
		PidInfoList []*PidInfo `json:"pidInfoList"`
		Total       int        `json:"total"`
	} `json:"result"`
}
