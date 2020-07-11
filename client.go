package vipsdk

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/varluffy/greq"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	ProdOpenAPIURL = "https://gw.vipapis.com"
	DefaultVersion = "1.0.0"
)

type Client struct {
	appKey      string
	appSecret   string
	host        string
	accessToken string
	req         *greq.Request
	resp        *greq.Response
	err         error
}

func NewClient(appKey, appSecret, host string, token ...string) *Client {
	client := new(Client)
	client.appKey = appKey
	client.appSecret = appSecret
	client.host = host
	if len(token) > 0 {
		client.accessToken = token[0]
	}
	return client
}

func (a *Client) UpdateKey(appKey, appSecret string) {
	a.appKey = appKey
	a.appSecret = appSecret
}

func (a *Client) SetParam(param Param) *Client {
	p := make(url.Values)
	//var p url.Values
	p.Add("appKey", a.appKey)
	p.Add("format", "json")
	p.Add("method", param.MethodName())
	p.Add("service", param.ServiceName())
	p.Add("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	p.Add("version", param.Version())
	req := greq.NewRequest("post", a.host)
	ps := param.Params()
	req.SetBodyJSON(ps)
	b, err := json.Marshal(ps)
	if err != nil {
		a.err = err
		return a
	}
	fmt.Println("body", string(b))
	p.Add("sign", a.sign(p, string(b)))
	if param.Token() {
		p.Add("accessToken", a.accessToken)
	}
	req.SetParams(p)
	a.req = req
	return a
}

func (a *Client) Exec(out interface{}) error {
	if a.err != nil {
		return a.err
	}
	if a.req == nil {
		return errors.New("a.req cannot be nil")
	}
	resp := a.req.Exec()
	if err := resp.Error(); err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("excepted statusCode=200 actual statusCode=%d", resp.StatusCode())
	}
	dumpRequest, _ := resp.DumpRequest(true)
	fmt.Println("dumpRequest", string(dumpRequest))
	if err := resp.ToJSON(out); err != nil {
		return err
	}
	return nil
}

func (a *Client) sign(param url.Values, body string) string {
	var strArr []string
	for key, _ := range param {
		strArr = append(strArr, key)
	}
	sort.Strings(strArr)
	var str string
	for _, key := range strArr {
		str = fmt.Sprintf("%s%s%v", str, key, param[key][0])
	}
	str = str + body
	h := hmac.New(md5.New, []byte(a.appSecret))
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
