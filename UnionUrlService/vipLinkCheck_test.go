package UnionUrlService

import "testing"

func TestVipLinkCheck(t *testing.T) {
	var p = VipLinkCheckParam {
		Request: VipLinkCheckReq{
			Content: "https://t.vip.com/wvDXpjGPrh7",
			Source: "yqd",
		},
	}

	var out VipLinkCheckResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}

	t.Logf("out: %+v", out)
	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
}
