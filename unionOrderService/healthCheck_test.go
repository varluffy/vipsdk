package unionOrderService

import (
	"github.com/varluffy/vipsdk"
	"testing"
)


var client = vipsdk.NewClient("", "", vipsdk.ProdOpenAPIURL)

func TestHealthCheck(t *testing.T) {
	var p = HealthCheckParam{}
	var out CheckResult
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}
	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
	t.Logf("%+v", out)
}
