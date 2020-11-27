package lubantop

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/iszhange/lubantop/requests"
)

const (
	HTTP_GATEWAY           = "http://gw.api.taobao.com/router/rest"
	HTTPS_GATEWAY          = "https://eco.taobao.com/router/rest"
	HTTP_TEST_GATEWAY      = "http://gw.api.tbsandbox.com/router/rest"
	HTTPS_TEST_GATEWAY     = "https://gw.api.tbsandbox.com/router/rest"
	HTTP_OVERSEAS_GATEWAY  = "http://api.taobao.com/router/rest"
	HTTPS_OVERSEAS_GATEWAY = "https://api.taobao.com/router/rest"
	MD5                    = "md5"
	HMAC                   = "hmac"
	XML                    = "xml"
	JSON                   = "json"
)

type Client struct {
	Gateway      string
	AppKey       string // TOP分配给应用的AppKey
	AppSecret    string // TOP分配给应用的AppSecret
	TargetAppKey string // 被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效
	SignMethod   string // 签名的摘要算法，可选值为：hmac，md5
	Format       string // 响应格式。默认为xml格式，可选值：xml，json。
	V            string // API协议版本，可选值：2.0
	Session      string // 用户登录授权成功后，TOP颁发给应用的授权信息，详细介绍请点击这里。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
	PartnerId    string // 合作伙伴身份标识
	Timeout      int64  // 请求超时时间
	method       string // API接口名称
	sign         string
	timestamp    string // 时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8，例如：2015-01-01 12:00:00。淘宝API服务端允许客户端请求最大时间误差为10分钟
	simplify     bool   // 是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false
	req          requests.Request
}

func New() *Client {
	c := new(Client)
	c.Gateway = HTTPS_GATEWAY
	c.SignMethod = MD5
	c.Format = JSON
	c.simplify = true
	c.V = "2.0"
	c.Timeout = 3

	return c
}

func (c *Client) Exec(req requests.Request) ([]byte, error) {
	c.req = req
	params := c.getParams()
	params["sign"] = c.getSign(params)

	return c.post(params)
}

func (c *Client) getParams() requests.Parameter {
	params := c.req.GetApiParams()
	params["method"] = c.req.GetMethod()
	params["app_key"] = c.AppKey
	params["sign_method"] = c.SignMethod
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["format"] = c.Format
	params["v"] = c.V
	if c.TargetAppKey != "" {
		params["target_app_key"] = c.TargetAppKey
	}
	if c.Session != "" {
		params["session"] = c.Session
	}
	if c.PartnerId != "" {
		params["partner_id"] = c.PartnerId
	}
	if c.simplify {
		params["simplify"] = c.simplify
	}

	return params
}

// 获取签名
func (c *Client) getSign(params requests.Parameter) string {
	// 获取Key
	var keys []string
	delete(params, "sign")
	for k := range params {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	query := bytes.NewBufferString(c.AppSecret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(interfaceToString(params[k]))
	}
	query.WriteString(c.AppSecret)
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, query)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (c *Client) post(params requests.Parameter) ([]byte, error) {
	var req *http.Request
	req, err := http.NewRequest("POST", c.Gateway, strings.NewReader(buildQuery(params)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpClient := &http.Client{}
	httpClient.Timeout = time.Second * time.Duration(c.Timeout)
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("请求错误:%d", response.StatusCode))
	}

	return ioutil.ReadAll(response.Body)
}

func buildQuery(param requests.Parameter) string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range param {
		args.Set(key, interfaceToString(val))
	}
	return args.Encode()
}

func interfaceToString(src interface{}) string {
	if src == nil {
		panic(requests.NIL_TYPE_ERROR)
	}

	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}

	return string(data)
}
