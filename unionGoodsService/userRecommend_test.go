package unionGoodsService

import "testing"

func TestUserRecommend(t *testing.T) {
	var p = UserRecommendParam{
		Request: UserRecommendRequest{
			Page:      1,
			RequestId: "123123123",
		},
	}

	var out UserRecommendResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}
	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
	t.Logf("out: %+v", out)
	for _, v := range out.Result.GoodsInfoList {
		t.Logf("goodsInfo: %+v", v)
	}
	for _, v := range out.Result.SortFields {
		t.Logf("SortField: %+v", v)
	}
}
