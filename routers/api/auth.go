package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/YeLlowaine/YeLlow/models"
	"github.com/YeLlowaine/YeLlow/pkg/e"
	"github.com/YeLlowaine/YeLlow/pkg/logging"
	"github.com/YeLlowaine/YeLlow/pkg/util"
)

type auth struct {
	Username         string `valid:"Required; MaxSize(50)"`
	Password         string `valid:"Required; MaxSize(50)"`
	FacePicture      string
	Icon             string
	Address          string
	SecurityQuestion string
	Duration         string
	UserType         int
}

//GetAuth ...
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//GetRegister
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user_type := com.StrTo(c.Query("user_type")).MustInt()
	face_picture := c.Query("face_picture")
	icon := c.Query("icon")
	address := c.Query("address")
	security_question := c.Query("security_question")
	duration := c.Query("duration")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("名称不能为空")
	valid.MaxSize(username, 100, "username").Message("名称最长为100字符")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MaxSize(password, 100, "password").Message("密码最长为100字符")
	valid.Range(user_type, 0, 1, "user_type").Message("用户种类只允许0或1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		models.AddAuth(username, password, face_picture, icon, address, security_question, duration, user_type)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func CheckAnswer(c *gin.Context) {
	username := c.Query("username")
	security_question := c.Query("security_question")

	valid := validation.Validation{}
	valid.Required(security_question, "security_question").Message("密保不能为空")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		isExist := models.CheckAnswer(username, security_question)
		if isExist {
			code = e.SUCCESS
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func UpdateAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	valid.Required(password, "password").Message("密码不能为空")
	valid.MaxSize(password, 100, "password").Message("密码最长为100字符")

	data := make(map[string]interface{})
	data["password"] = password
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		models.UpdateAuth(username, data)
		code = e.SUCCESS
	} else {
		code = e.ERROR
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func SearchUser(c *gin.Context) {
	input := c.Query("keyword")
	data := make(map[string]interface{})
	keyword := "%" + input + "%"
	valid := validation.Validation{}
	valid.Required(keyword, "keyword").Message("关键词不能为空")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.SearchUser(keyword)
		for index, val := range ret {
			str := fmt.Sprintf("%d", index)
			data[str] = val
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// //ChangeAuth ...
// func ChangeAuth(c *gin.Context) {
// 	username := c.Query("username")
// 	password := c.Query("password")

// 	valid := validation.Validation{}
// 	a := auth{Username: username, Password: password}
// 	ok, _ := valid.Valid(&a)

// 	data := make(map[string]interface{})
// 	code := e.INVALID_PARAMS
// 	if ok {
// 		isExist := models.CheckAuth(username, password)
// 		if isExist {
// 			token, err := util.GenerateToken(username, password)
// 			if err != nil {
// 				code = e.ERROR_AUTH_TOKEN
// 			} else {
// 				data["token"] = token

// 				code = e.SUCCESS
// 			}

// 		} else {
// 			code = e.ERROR_AUTH
// 		}
// 	} else {
// 		for _, err := range valid.Errors {
// 			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": code,
// 		"msg":  e.GetMsg(code),
// 		"data": data,
// 	})
// }

//GetInfo 获取用户信息
func GetInfo(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}

	var data interface{}
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		data = models.GetInfo(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//GetInfoFromName 通过用户名取用户信息
func GetInfoFromName(c *gin.Context) {
	username := c.Param("username")

	valid := validation.Validation{}

	var data interface{}
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		data = models.GetInfoFromName(username)
		code = e.SUCCESS
	} else {
		code = e.ERROR
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//GetIcon 通过用户名取用户头像
func GetIcon(c *gin.Context) {
	username := c.Param("username")

	valid := validation.Validation{}

	var data interface{}
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		data = models.GetIcon(username)
		code = e.SUCCESS
	} else {
		code = e.ERROR
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetAllDoctor(c *gin.Context) {

	data := make(map[string]interface{})
	ret := models.GetAll()
	code := e.SUCCESS
	for index, val := range ret {
		str := fmt.Sprintf("%d", index)
		data[str] = val
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
