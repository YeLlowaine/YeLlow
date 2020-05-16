package im

import (
	"log"
	"net/http"

	"github.com/YeLlowaine/YeLlow/pkg/e"
	irc "github.com/YeLlowaine/YeLlow/pkg/rc"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

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
func PrivateSend(c *gin.Context) {
	senderID := c.Query("senderID")
	targetIDs := make([]string, 10)
	targetID := append(targetIDs, c.Query("targetID"))
	objectName := c.Query("objectName")

	msg := c.Query("msg")
	pushContent := c.Query("pushContent")
	pushData := c.Query("pushData")
	count := com.StrTo(c.Query("count")).MustInt()
	verifyBlacklist := com.StrTo(c.DefaultQuery("verifyBlacklist", "0")).MustInt()
	isPersisted := com.StrTo(c.DefaultQuery("isPersisted", "0")).MustInt()
	isIncludeSender := com.StrTo(c.DefaultQuery("isIncludeSender", "0")).MustInt()
	contentAvailable := com.StrTo(c.Query("contentAvailable")).MustInt()
	appKey := "k51hidwqkvvqb"
	appSec := "2LbRng4DyGoYP"
	rc := irc.NewRongCloud(
		appKey,
		appSec,
	)

	err := rc.PrivateSend(senderID, targetID, objectName, msg, pushContent, pushData, count, verifyBlacklist, isPersisted, isIncludeSender, contentAvailable)

	code := e.INVALID_PARAMS
	if err != nil {
		log.Print(err)
		code = e.ERROR
	} else {
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// PrivateSendRecall 撤回单聊消息方法
/*
*
*@param  senderID:发送人用户 ID。
*@param  targetID:接收用户 ID。
*@param  uID:消息的唯一标识，各端 v1 发送消息成功后会返回 uID。
*@param  sentTime:消息的发送时间，各端 v1 发送消息成功后会返回 sentTime。
*
*@return error
 */
func PrivateSendRecall(c *gin.Context) {
	senderID := c.Query("senderID")
	targetID := c.Query("targetID")
	uID := c.Query("uID")

	sentTime := com.StrTo(c.Query("sentTime")).MustInt()
	appKey := "k51hidwqkvvqb"
	appSec := "2LbRng4DyGoYP"
	rc := irc.NewRongCloud(
		appKey,
		appSec,
	)

	err := rc.PrivateRecall(senderID, targetID, uID, sentTime)

	code := e.INVALID_PARAMS
	if err != nil {
		log.Print(err)
		code = e.ERROR
	} else {
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
