package models

import (
	"compress/gzip"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego/httplib"
)

const (
	// RONGCLOUDSMSURI 融云默认 SMS API 地址
	RONGCLOUDSMSURI = "http://api.sms.ronghub.com"
	// RONGCLOUDURI 融云默认 API 地址
	RONGCLOUDURI = "http://api-cn.ronghub.com"
	// RONGCLOUDURI2 融云备用 API 地址
	RONGCLOUDURI2 = "http://api2-cn.ronghub.com"
	// ReqType body类型
	ReqType = "json"
	// USERAGENT sdk 名称
	USERAGENT = "rc-go-sdk/3.0.2"
	// DEFAULTTIMEOUT 默认超时时间
	DEFAULTTIMEOUT = 10
	// NUMTIMEOUT 默认超时次数切换 Api 地址
	NUMTIMEOUT = 3
)

var (
	defaultExtra = rongCloudExtra{
		rongCloudURI:    RONGCLOUDURI,
		rongCloudSMSURI: RONGCLOUDSMSURI,
		timeout:         DEFAULTTIMEOUT,
		numTimeout:      NUMTIMEOUT,
		count:           0,
	}
	rc   *RongCloud
	once sync.Once
)

// RongCloud appKey appSecret extra
type RongCloud struct {
	appKey    string
	appSecret string
	*rongCloudExtra
}

// CodeResult 融云返回状态码和错误码
type CodeResult struct {
	Code         int    `json:"code"`         // 返回码，200 为正常。
	ErrorMessage string `json:"errorMessage"` // 错误信息
}

// rongCloudExtra rongCloud扩展增加自定义融云服务器地址,请求超时时间
type rongCloudExtra struct {
	rongCloudURI    string
	rongCloudSMSURI string
	timeout         time.Duration
	numTimeout      uint
	count           uint
}

// getSignature 本地生成签名
// Signature (数据签名)计算方法：将系统分配的 App Secret、Nonce (随机数)、
// Timestamp (时间戳)三个字符串按先后顺序拼接成一个字符串并进行 SHA1 哈希计算。如果调用的数据签名验证失败，接口调用会返回 HTTP 状态码 401。
func (rc RongCloud) getSignature() (nonce, timestamp, signature string) {
	nonceInt := rand.Int()
	nonce = strconv.Itoa(nonceInt)
	timeInt64 := time.Now().Unix()
	timestamp = strconv.FormatInt(timeInt64, 10)
	h := sha1.New()
	_, _ = io.WriteString(h, rc.appSecret+nonce+timestamp)
	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

// fillHeader 在 Http Header 增加API签名
func (rc RongCloud) fillHeader(req *httplib.BeegoHTTPRequest) {
	nonce, timestamp, signature := rc.getSignature()
	req.Header("App-Key", rc.appKey)
	req.Header("Nonce", nonce)
	req.Header("Timestamp", timestamp)
	req.Header("Signature", signature)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Header("User-Agent", USERAGENT)
}

// fillJSONHeader 在 Http Header Content-Type 设置为josn格式
func fillJSONHeader(req *httplib.BeegoHTTPRequest) {
	req.Header("Content-Type", "application/json")
}

// GetRongCloud 获取 RongCloud 对象
func GetRongCloud() *RongCloud {
	return rc
}

// changeURI 切换 Api 服务器地址
func (rc *RongCloud) changeURI() {
	switch rc.rongCloudURI {
	case RONGCLOUDURI:
		rc.rongCloudURI = RONGCLOUDURI2
	case RONGCLOUDURI2:
		rc.rongCloudURI = RONGCLOUDURI
	default:
	}
}

// PrivateURI 私有云设置 Api 地址
func (rc *RongCloud) PrivateURI(uri, sms string) {
	rc.rongCloudURI = uri
	rc.rongCloudSMSURI = sms
}

/**
判断 http status code, 如果大于 500 就切换一次域名
*/
func (rc *RongCloud) checkStatusCode(resp *http.Response) {
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		rc.changeURI()
	}

	return
}

// rcMsg rcMsg接口
type rcMsg interface {
	ToString() (string, error)
}

// MsgUserInfo 融云内置消息用户信息
type MsgUserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Extra string `json:"extra"`
}

// TXTMsg 消息
type TXTMsg struct {
	Content string      `json:"content"`
	User    MsgUserInfo `json:"user"`
	Extra   string      `json:"extra"`
}

type rongCloudOption func(*RongCloud)

// ToString TXTMsg
func (msg *TXTMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func NewRongCloud(appKey, appSecret string) *RongCloud {
	once.Do(func() {
		// 默认扩展配置
		defaultRongCloud := defaultExtra
		rc = &RongCloud{
			appKey:         appKey,
			appSecret:      appSecret,
			rongCloudExtra: &defaultRongCloud,
		}
	},
	)

	return rc
}

// RCErrorNew 创建新的err信息
func RCErrorNew(code int, text string) error {
	return CodeResult{code, text}
}

// Error 获取错误信息
func (e CodeResult) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.ErrorMessage
}

func (rc *RongCloud) do(b *httplib.BeegoHTTPRequest) (body []byte, err error) {
	return rc.httpRequest(b)
}

func (rc *RongCloud) httpRequest(b *httplib.BeegoHTTPRequest) (body []byte, err error) {
	resp, err := b.DoRequest()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	rc.checkStatusCode(resp)
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err = checkHTTPResponseCode(body); err != nil {
		return nil, err
	}
	return body, err
}

func checkHTTPResponseCode(rep []byte) error {
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return code
	}
	return nil
}

func (rc *RongCloud) PrivateSend(senderID string, targetID []string, objectName string, msg rcMsg,
	pushContent, pushData string, count, verifyBlacklist, isPersisted, isIncludeSender, contentAvailable int) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	if len(targetID) == 0 {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(RONGCLOUDURI + "/message/private/publish." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("fromUserId", senderID)
	for _, v := range targetID {
		req.Param("toUserId", v)
	}
	req.Param("objectName", objectName)

	msgr, err := msg.ToString()
	if err != nil {
		return err
	}
	req.Param("content", msgr)
	req.Param("pushData", pushData)
	req.Param("pushContent", pushContent)
	req.Param("count", strconv.Itoa(count))
	req.Param("verifyBlacklist", strconv.Itoa(verifyBlacklist))
	req.Param("isPersisted", strconv.Itoa(isPersisted))
	req.Param("contentAvailable", strconv.Itoa(contentAvailable))
	req.Param("isIncludeSender", strconv.Itoa(isIncludeSender))

	return err
}
