package unionUrlService

import (
	"github.com/varluffy/vipsdk"
	"testing"
)

var client = vipsdk.NewClient("", "", vipsdk.ProdOpenAPIURL)

func TestGenByGoodsId(t *testing.T) {
	var p = GenByGoodsIdParam{
		Request: GenByGoodsIdRequest{
			RequestId:   "12312312333333",
			GoodsIdList: []string{"6918712886710583375"},
			ChanTag:     "12312312321321",
		},
	}

	var out UrlGenResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}

	t.Logf("out: %+v", out)
	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode+out.ReturnMessage)
		return
	}
	for _, v := range out.Result.URLInfoList {
		t.Logf("result: %+v", v)
	}
}
