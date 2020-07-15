package unionPidService

import "testing"

func TestQueryPid(t *testing.T) {
	var p = QueryPidParam{
		Request: PidQueryRequest{
			RequestId: "123123123",
		},
	}

	var out CpsUnionPidQueryResponse
	err := client.SetParam(p).Exec(&out)
	if err != nil {
		t.Error(err)
	}

	t.Logf("out: %+v", out)
	if out.ReturnCode != "0" {
		t.Errorf("return code not equal 0, code: %s", out.ReturnCode)
		return
	}
	for _, v := range out.Result.PidInfoList {
		t.Logf("result: %+v", v)
	}
}
