package unionGoodsService

import (
	"github.com/varluffy/vipsdk"
	"testing"
)

var client = vipsdk.NewClient("", "", vipsdk.ProdOpenAPIURL)

func TestGetByGoodsIds(t *testing.T) {
	var p = GetByGoodsIdsParam{
		Request: GetByGoodsIdsRequest{
			GoodsIdList: []string{"6918733344432093713"},
			RequestId:   "requestid",
		},
	}
	var out GetByGoodsIdsResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", out.Result)
	for _, v := range out.Result {
		t.Logf("out: %+v", v)
	}

}
