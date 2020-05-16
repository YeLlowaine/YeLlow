package im

import (
	"log"
	"net/http"

	"github.com/YeLlowaine/YeLlow/pkg/e"
	irc "github.com/YeLlowaine/YeLlow/pkg/rc"
	"github.com/gin-gonic/gin"
)

// GetToken 注册用户，生成用户在融云的唯一身份标识 Token
/*
*@param  userID:用户 ID，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*@param  name:用户名称，最大长度 128 字节.用来在 Push 推送时显示用户的名称.用户名称，最大长度 128 字节.用来在 Push 推送时显示用户的名称。
*@param  portraitURI:用户头像 URI，最大长度 1024 字节.用来在 Push 推送时显示用户的头像。
*
*@return User, error
 */
func GetToken(c *gin.Context) {
	userID := c.Query("userID")
	name := c.Query("name")
	potraitURI := c.Query("potraitURI")
	appKey := "k51hidwqkvvqb"
	appSec := "2LbRng4DyGoYP"
	rc := irc.NewRongCloud(
		appKey,
		appSec,
	)

	token, err := rc.UserRegister(userID, name, potraitURI)
	log.Print(token.Token)
	code := e.INVALID_PARAMS
	if err != nil {
		log.Print(err)
		code = e.ERROR
	} else {
		code = e.SUCCESS
	}
	data := make(map[string]string)
	data["token"] = token.Token
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// UserUpdate 修改用户信息
/*
*@param  userID:用户 ID，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*@param  name:用户名称，最大长度 128 字节。用来在 Push 推送时，显示用户的名称，刷新用户名称后 5 分钟内生效。（可选，提供即刷新，不提供忽略）
*@param  portraitURI:用户头像 URI，最大长度 1024 字节。用来在 Push 推送时显示。（可选，提供即刷新，不提供忽略）
*
*@return error
 */
func UserUpdate(c *gin.Context) {
	userID := c.Query("userID")
	name := c.Query("name")
	potraitURI := c.Query("potraitURI")
	appKey := "k51hidwqkvvqb"
	appSec := "2LbRng4DyGoYP"
	rc := irc.NewRongCloud(
		appKey,
		appSec,
	)

	err := rc.UserUpdate(userID, name, potraitURI)

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

// OnlineStatusCheck 检查用户在线状态
/*
*@param  userID:用户 ID，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*
*@return int, error
 */
func OnlineStatusCheck(c *gin.Context) {
	userID := c.Query("userID")
	appKey := "k51hidwqkvvqb"
	appSec := "2LbRng4DyGoYP"
	rc := irc.NewRongCloud(
		appKey,
		appSec,
	)

	status, err := rc.OnlineStatusCheck(userID)
	data := make(map[string]int)
	data["status"] = status
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
		"data": data,
	})
}
