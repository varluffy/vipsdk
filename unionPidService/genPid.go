package unionPidService

// 市场和渠道服务 / 联盟PID开放服务 / 创建推广位PID
// http://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidService-1.0.0/genPid
type GenPidParam struct {
	Service
	Request PidGenRequest
}

func (a GenPidParam) MethodName() string {
	return "genPid"
}

func (a GenPidParam) Params() interface{} {
	p := make(map[string]interface{}, 0)
	p["pidGenRequest"] = a.Request
	return p
}

// 推广位创建请求信息
type PidGenRequest struct {
	PidNameList []string `json:"pidNameList"` // 需要生成的推广位名称列表 注： 1、一次支持批量最大100个 2、每个推广位的名称最长50个字符
	RequestId   string   `json:"requestId"`   // 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
}

type UnionPidGenResponse struct {
	ReturnCode string `json:"returnCode"`
	Result     struct {
		PidInfoList    []*PidInfo `json:"pidInfoList"`
		Total          int        `json:"total"`
		RemainPidCount int        `json:"remainPidCount"`
	} `json:"result"`
}
