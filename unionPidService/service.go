package unionPidService

import "github.com/varluffy/vipsdk"

type Service struct{}

func (a Service) ServiceName() string {
	return "com.vip.adp.api.open.service.UnionPidService"
}

func (a Service) Version() string {
	return vipsdk.DefaultVersion
}

func (a Service) Token() bool {
	return false
}

// 推广位Pid信息列表
type PidInfo struct {
	Pid        string `json:"pid"` // 推广位ID
	PidName    string `json:"pidName"` // 推广位名称
	CreateTime int    `json:"createTime"` // 该推广位创建的时间
	Ucode      string `json:"ucode"` // 所属ucode
	B2CUserID  string `json:"b2cUserId"` // 所属b2c_user_id
	Status     int    `json:"status"` // pid状态，1表示可用，2表示禁用
}