package unionOrderService

type HealthCheckParam struct {
	Service
}

func (a HealthCheckParam) MethodName() string {
	return "healthCheck"
}

func (a HealthCheckParam) Params() interface{} {
	return ""
}

type CheckResult struct {
	ReturnCode string `json:"returnCode"`
	Result     struct {
		Status  Status `json:"status"`
		Details map[string]string
	} `json:"result"`
}

type Status struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
