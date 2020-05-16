package v1

import (
	// "log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/YeLlowaine/YeLlow/models"
	"github.com/YeLlowaine/YeLlow/pkg/e"
	"github.com/YeLlowaine/YeLlow/pkg/logging"
	"github.com/YeLlowaine/YeLlow/pkg/setting"
	"github.com/YeLlowaine/YeLlow/pkg/util"
)

//GetComment 获取单个评论
func GetComment(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistCommentByID(id) {
			data = models.GetComment(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_COMMENT
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

//GetComments 获取多个评论
func GetComments(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var articleId int = -1
	if arg := c.Query("article_id"); arg != "" {
		articleId = com.StrTo(arg).MustInt()
		maps["article_id"] = articleId

		valid.Min(articleId, 1, "article_id").Message("文章ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetComments(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetCommentTotal(maps)

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

//AddComment 新增评论
func AddComment(c *gin.Context) {
	articleId := com.StrTo(c.Query("article_id")).MustInt()
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	

	valid := validation.Validation{}
	valid.Min(articleId, 1, "article_id").Message("文章ID必须大于0")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(articleId) {
			data := make(map[string]interface{})
			data["article_id"] = articleId
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddComment(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
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

// //DeleteComment 删除评论
// func DeleteComment(c *gin.Context) {
// 	id := com.StrTo(c.Param("id")).MustInt()

// 	valid := validation.Validation{}
// 	valid.Min(id, 1, "id").Message("ID必须大于0")

// 	code := e.INVALID_PARAMS
// 	if !valid.HasErrors() {
// 		if models.ExistCommentByID(id) {
// 			models.DeleteComment(id)
// 			code = e.SUCCESS
// 		} else {
// 			code = e.ERROR_NOT_EXIST_COMMENT
// 		}
// 	} else {
// 		for _, err := range valid.Errors {
// 			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": code,
// 		"msg":  e.GetMsg(code),
// 		"data": make(map[string]string),
// 	})
// }
