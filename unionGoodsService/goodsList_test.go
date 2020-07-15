package unionGoodsService

import "testing"

func TestGoodsList(t *testing.T) {
	var p = GoodsListParam{
		Request: GoodsInfoRequest{
			Page:      1,
			RequestId: "123123123",
		},
	}

	var out interface{}
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}
	t.Logf("out: %#v", out)
}
