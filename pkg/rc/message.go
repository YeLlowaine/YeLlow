/*
 * @Descripttion:
 * @version:
 * @Author: ran.ding
 * @Date: 2019-09-02 18:29:55
 * @LastEditors: ran.ding
 * @LastEditTime: 2019-09-10 11:37:14
 */
package rc

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

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

// ImgMsg 消息
type ImgMsg struct {
	Content  string      `json:"content"`
	User     MsgUserInfo `json:"user"`
	ImageURI string      `json:"imageUri"`
	Extra    string      `json:"extra"`
}

// InfoNtf 消息
type InfoNtf struct {
	Message string      `json:"message"`
	User    MsgUserInfo `json:"user"`
	Extra   string      `json:"extra"`
}

// VCMsg 消息
type VCMsg struct {
	Content  string      `json:"content"`
	User     MsgUserInfo `json:"user"`
	Extra    string      `json:"extra"`
	Duration string      `json:"duration"`
}

// IMGTextMsg 消息
type IMGTextMsg struct {
	Title    string      `json:"title"`
	Content  string      `json:"content"`
	User     MsgUserInfo `json:"user"`
	Extra    string      `json:"extra"`
	Duration string      `json:"duration"`
	URL      string      `json:"url"`
}

// FileMsg 消息
type FileMsg struct {
	Name    string      `json:"name"`
	Size    string      `json:"size"`
	Type    string      `json:"type"`
	FileURL string      `json:"fileUrl"`
	User    MsgUserInfo `json:"user"`
}

// LBSMsg 消息
type LBSMsg struct {
	Content   string      `json:"content"`
	Extra     string      `json:"extra"`
	POI       string      `json:"poi"`
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	User      MsgUserInfo `json:"user"`
}

// ProfileNtf 消息
type ProfileNtf struct {
	Operation string      `json:"operation"`
	Data      string      `json:"data"`
	User      MsgUserInfo `json:"user"`
	Extra     string      `json:"extra"`
}

// CMDNtf 消息
type CMDNtf struct {
	Name string      `json:"operation"`
	Data string      `json:"data"`
	User MsgUserInfo `json:"user"`
}

// CMDMsg 消息
type CMDMsg struct {
	Name string      `json:"operation"`
	Data string      `json:"data"`
	User MsgUserInfo `json:"user"`
}

// ContactNtf 消息
type ContactNtf struct {
	Operation    string      `json:"operation"`
	SourceUserID string      `json:"sourceUserId"`
	TargetUserID string      `json:"targetUserId"`
	Message      string      `json:"message"`
	Extra        string      `json:"extra"`
	User         MsgUserInfo `json:"user"`
}

// GrpNtf 消息
type GrpNtf struct {
	OperatorUserID string      `json:"operatorUserId"`
	Operation      string      `json:"operation"`
	Data           string      `json:"data"`
	Message        string      `json:"message"`
	Extra          string      `json:"extra"`
	User           MsgUserInfo `json:"user"`
}

// DizNtf 消息
type DizNtf struct {
	Type      int         `json:"type"`
	Extension string      `json:"extension"`
	Operation string      `json:"operation"`
	User      MsgUserInfo `json:"user"`
}

// TemplateMsgContent 消息模版
type TemplateMsgContent struct {
	TargetID    string
	Data        map[string]string
	PushContent string
}

// MentionedInfo Mentioned
type MentionedInfo struct {
	Type        int      `json:"type"`
	UserIDs     []string `json:"userIdList"`
	PushContent string   `json:"mentionedContent"`
}

// MentionMsgContent MentionMsgContent
type MentionMsgContent struct {
	Content       string        `json:"content"`
	MentionedInfo MentionedInfo `json:"mentionedinfo"`
}

// History History
type History struct {
	URL string `json:"url"`
}

// BroadcastRecallContent content of message broadcast recall
type BroadcastRecallContent struct {
	MessageId        string `json:"messageUId"`
	ConversationType int    `json:"conversationType"`
	IsAdmin          int    `json:"isAdmin"`
	IsDelete         int    `json:"isDelete"`
}

/**
 * @name: ToString
 * @test:
 * @msg: 将广播消息撤回接口需要的 content 结构体参数转换为 json
 * @param {type}
 * @return: string error
 */
func (content *BroadcastRecallContent) ToString() (string, error) {
	bytes, err := json.Marshal(content)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// ToString TXTMsg
func (msg *TXTMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString ImgMsg
func (msg *ImgMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString InfoNtf
func (msg *InfoNtf) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString VCMsg
func (msg *VCMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString IMGTextMsg
func (msg *IMGTextMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString FileMsg
func (msg *FileMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString LBSMsg
func (msg *LBSMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString ProfileNtf
func (msg *ProfileNtf) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString CMDNtf
func (msg *CMDNtf) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString CMDMsg
func (msg *CMDMsg) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString ContactNtf
func (msg *ContactNtf) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString GrpNtf
func (msg *GrpNtf) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToString DizNtf
func (msg *DizNtf) ToString() (string, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// PrivateSend 发送单聊消息方法（一个用户向多个用户发送消息，单条消息最大 128k。每分钟最多发送 6000 条信息，每次发送用户上限为 1000 人，如：一次发送 1000 人时，示为 1000 条消息。）
/*
 *@param  senderID:发送人用户 ID。
 *@param  targetID:接收用户 ID。可以实现向多人发送消息，每次上限为 1000 人。
 *@param  objectName:发送的消息类型。
 *@param  msg:消息内容。
 *@param  pushContent:定义显示的 Push 内容，如果 objectName 为融云内置消息类型时，则发送后用户一定会收到 Push 信息。如果为自定义消息，则 pushContent 为自定义消息显示的 Push 内容，如果不传则用户不会收到 Push 通知。
 *@param  pushData:针对 iOS 平台为 Push 通知时附加到 payload 中，Android 客户端收到推送消息时对应字段名为 pushData。
 *@param  count:针对 iOS 平台，Push 时用来控制未读消息显示数，只有在 toUserId 为一个用户 Id 的时候有效。
 *@param  verifyBlacklist:是否过滤发送人黑名单列表，0 表示为不过滤、 1 表示为过滤，默认为 0 不过滤。
 *@param  isPersisted:当前版本有新的自定义消息，而老版本没有该自定义消息时，老版本客户端收到消息后是否进行存储，0 表示为不存储、 1 表示为存储，默认为 1 存储消息。
 *@param  isIncludeSender:发送用户自已是否接收消息，0 表示为不接收，1 表示为接收，默认为 0 不接收。
 *@param  contentAvailable:针对 iOS 平台，对 v1 处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行，查看详细。1 表示为开启，0 表示为关闭，默认为 0。
 *
 *@return error
 */

func (rc *RongCloud) PrivateSend(senderID string, targetID []string, objectName string, msg string,
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

	req.Param("content", msg)
	req.Param("pushData", pushData)
	req.Param("pushContent", pushContent)
	req.Param("count", strconv.Itoa(count))
	req.Param("verifyBlacklist", strconv.Itoa(verifyBlacklist))
	req.Param("isPersisted", strconv.Itoa(isPersisted))
	req.Param("contentAvailable", strconv.Itoa(contentAvailable))
	req.Param("isIncludeSender", strconv.Itoa(isIncludeSender))

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// PrivateRecall 撤回单聊消息方法
/*
*
*@param  senderID:发送人用户 ID。
*@param  targetID:接收用户 ID。
*@param  uID:消息的唯一标识，各端 v1 发送消息成功后会返回 uID。
*@param  sentTime:消息的发送时间，各端 v1 发送消息成功后会返回 sentTime。
*
*@return error
 */

func (rc *RongCloud) PrivateRecall(senderID, targetID, uID string, sentTime int) error {
	if senderID == "" {
		return RCErrorNew(1002, "Paramer 'senderID' is required")
	}

	if targetID == "" {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/message/recall." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("fromUserId", senderID)
	req.Param("targetId", targetID)
	req.Param("messageUID", uID)
	req.Param("sentTime", strconv.Itoa(sentTime))
	req.Param("conversationType", strconv.Itoa(1))

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}
