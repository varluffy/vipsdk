package unionUrlService

import (
	"github.com/varluffy/vipsdk"
	"testing"
)

var client = vipsdk.NewClient("0b78bc17", "CAC03E0A28BA1539802A2D7A312F4345", vipsdk.ProdOpenAPIURL)

func TestGenByGoodsId(t *testing.T) {
	var p = GenByGoodsIdParam{
		Request: GenByGoodsIdRequest{
			RequestId:   "123123123",
			GoodsIdList: []string{"6918733344432093713"},
		},
	}

	var out UrlGenResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}

	t.Logf("out: %+v", out)
	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
	for _, v := range out.Result.URLInfoList {
		t.Logf("result: %+v", v)
	}
}
