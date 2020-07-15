package unionPidService

import (
	"github.com/varluffy/vipsdk"
	"testing"
)

var client = vipsdk.NewClient("", "", vipsdk.ProdOpenAPIURL)

func TestGenPid(t *testing.T) {
	var p = GenPidParam{
		Request: PidGenRequest{
			RequestId:   "123123123",
			PidNameList: []string{"有钱得"},
		},
	}

	var out UnionPidGenResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}

	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
	t.Logf("out: %+v", out)
	for _, v := range out.Result.PidInfoList {
		t.Logf("result: %+v", v)
	}
}
