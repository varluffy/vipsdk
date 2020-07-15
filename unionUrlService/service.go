package unionUrlService

import "github.com/varluffy/vipsdk"

type Service struct{}

func (a Service) ServiceName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

func (a Service) Version() string {
	return vipsdk.DefaultVersion
}

func (a Service) Token() bool {
	return false
}
