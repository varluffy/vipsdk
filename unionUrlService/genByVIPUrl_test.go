package unionUrlService

import "testing"

func TestGenByVIPUrl(t *testing.T) {
	var p = GenByVIPUrlParam{
		Request: GenByVIPUrlRequest{
			RequestId: "123123123",
			UrlList:   []string{"https://m.vip.com/product-1710612817-6918733344432093713.html?wq=1", "https://m.vip.com/product-1711231439-6918723043060636495.html?nmsns=shop_iphone-7.24.3-link&nst=product&nsbc=&nct=link&ncid=0f4a0d096a9f92db7249f156e6ca43fd538c5665&nabtid=13&nuid=&nchl_param=share:0f4a0d096a9f92db7249f156e6ca43fd538c5665:1594435060511"},
		},
	}

	var out GenByVIPUrlResponse
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
