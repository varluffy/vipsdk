package unionOrderService

import "testing"

func TestRefundOrderList(t *testing.T)  {
	var p = RefundOrderListParam{
		Request: RefundOrderRequest{
			Page:1,
			RequestId: "123123123",
		},
	}

	var out RefundOrderResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}

	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
	t.Logf("out: %+v", out)
	for _, v := range out.Result.RefundOrderInfoList {
		t.Logf("result: %+v", v)
	}
}
