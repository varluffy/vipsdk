package unionOrderService

import "testing"

func TestOrderList(t *testing.T) {
	var p = OrderListParam{
		Request: OrderQueryModel{
			Page:      1,
			RequestId: "12312ssss3123",
			//OrderTimeStart: 1594447420652,
			//OrderTimeEnd:   1594447427652,
			UpdateTimeStart: 1595221738000,
			UpdateTimeEnd:   1595332738000,
			ChanTag: "[\"c21489ea1cb7e53356f88add1295f303\"]",
		},
	}

	var out OrderListResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}

	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
	t.Logf("out: %+v", out)
	for _, v := range out.Result.OrderInfoList {
		t.Logf("result: %+v", v)
	}
}
