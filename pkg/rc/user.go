/*
 * @Descripttion:
 * @version:
 * @Author: ran.ding
 * @Date: 2019-09-02 18:29:55
 * @LastEditors: ran.ding
 * @LastEditTime: 2019-09-10 11:22:38
 */
package rc

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

// User 用户信息 返回信息
type User struct {
	Token        string `json:"token"`
	UserID       string `json:"userId"`
	BlockEndTime string `json:"blockEndTime,omitempty"`
	Status       string `json:"status,omitempty"`
}

// UserRegister 注册用户，生成用户在融云的唯一身份标识 Token
/*
*@param  userID:用户 ID，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*@param  name:用户名称，最大长度 128 字节.用来在 Push 推送时显示用户的名称.用户名称，最大长度 128 字节.用来在 Push 推送时显示用户的名称。
*@param  portraitURI:用户头像 URI，最大长度 1024 字节.用来在 Push 推送时显示用户的头像。
*
*@return User, error
 */

func (rc *RongCloud) UserRegister(userID, name, portraitURI string) (User, error) {
	if userID == "" {
		return User{}, RCErrorNew(1002, "Paramer 'userID' is required")
	}
	if name == "" {
		return User{}, RCErrorNew(1002, "Paramer 'name' is required")
	}
	if portraitURI == "" {
		return User{}, RCErrorNew(1002, "Paramer 'portraitUri' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/user/getToken." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("userId", userID)
	req.Param("name", name)
	req.Param("portraitURI", portraitURI)

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return User{}, err
	}

	var userResult User
	if err := json.Unmarshal(resp, &userResult); err != nil {
		return User{}, err
	}
	return userResult, nil
}

// UserUpdate 修改用户信息
/*
*@param  userID:用户 ID，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*@param  name:用户名称，最大长度 128 字节。用来在 Push 推送时，显示用户的名称，刷新用户名称后 5 分钟内生效。（可选，提供即刷新，不提供忽略）
*@param  portraitURI:用户头像 URI，最大长度 1024 字节。用来在 Push 推送时显示。（可选，提供即刷新，不提供忽略）
*
*@return error
 */

func (rc *RongCloud) UserUpdate(userID, name, portraitURI string) error {
	if userID == "" {
		return RCErrorNew(1002, "Paramer 'userID' is required")
	}
	if name == "" {
		return RCErrorNew(1002, "Paramer 'name' is required")
	}
	if portraitURI == "" {
		return RCErrorNew(1002, "Paramer 'portraitURI' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/user/refresh." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("userId", userID)
	req.Param("name", name)
	req.Param("portraitUri", portraitURI)

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// OnlineStatusCheck 检查用户在线状态
/*
*@param  userID:用户 ID，最大长度 64 字节.是用户在 App 中的唯一标识码，必须保证在同一个 App 内不重复，重复的用户 Id 将被当作是同一用户。
*
*@return int, error
 */
func (rc *RongCloud) OnlineStatusCheck(userID string) (int, error) {
	if userID == "" {
		return -1, RCErrorNew(1002, "Paramer 'userID' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/user/checkOnline." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("userId", userID)

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return -1, err
	}
	var userResult User
	if err := json.Unmarshal(resp, &userResult); err != nil {
		return -1, err
	}
	status, _ := strconv.Atoi(userResult.Status)
	return status, nil
}
