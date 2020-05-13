package v1

import (
	"fmt"
	"net/http"

	"github.com/Hallelujah1025/Stroke-Survivors/models"
	"github.com/Hallelujah1025/Stroke-Survivors/pkg/e"
	"github.com/Hallelujah1025/Stroke-Survivors/pkg/logging"
	"github.com/Hallelujah1025/Stroke-Survivors/pkg/setting"
	"github.com/Hallelujah1025/Stroke-Survivors/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticles
func GetArticles(c *gin.Context) {

	article_name := c.Query("article_name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	if article_name != "" {
		maps["article_name"] = article_name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)

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

//GetArticlesFromSomeone 获取某人所有文章
func GetArticlesFromSomeone(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var createdBy string = c.Param("id")
	maps["created_by"] = createdBy

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)

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

//GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
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

// AddArticle
func AddArticle(c *gin.Context) {
	article_name := c.Query("article_name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	content := c.Query("content")
	valid := validation.Validation{}

	valid.Required(article_name, "article_name").Message("名称不能为空")
	valid.MaxSize(article_name, 100, "article_name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistArticleByName(article_name) {
			code = e.SUCCESS
			models.AddArticle(article_name, state, createdBy, content)
		} else {
			code = e.ERROR_EXIST_ARTICLE
		}
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

//DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
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
		"data": make(map[string]string),
	})
}

func SearchArticle(c *gin.Context) {
	input := c.Query("keyword")
	data := make(map[string]interface{})
	keyword := "%" + input + "%"
	valid := validation.Validation{}
	valid.Required(keyword, "keyword").Message("关键词不能为空")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.SearchArticle(keyword)
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

// // @Summary Update article
// // @Produce  json
// // @Param id path int true "ID"
// // @Param name body string true "ArticleName"
// // @Param state body int false "State"
// // @Param modified_by body string true "ModifiedBy"
// // @Success 200 {object} app.Response
// // @Failure 500 {object} app.Response
// // @Router /api/v1/tags/{id} [put]
// func EditArticle(c *gin.Context) {
// 	id := com.StrTo(c.Param("id")).MustInt()
// 	article_name := c.Query("article_name")
// 	modifiedBy := c.Query("modified_by")

// 	valid := validation.Validation{}

// 	var state int = -1
// 	if arg := c.Query("state"); arg != "" {
// 		state = com.StrTo(arg).MustInt()
// 		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
// 	}

// 	valid.Required(id, "id").Message("ID不能为空")
// 	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
// 	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
// 	valid.MaxSize(article_name, 100, "name").Message("名称最长为100字符")

// 	code := e.INVALID_PARAMS
// 	if !valid.HasErrors() {
// 		code = e.SUCCESS
// 		if models.ExistArticleByID(id) {
// 			data := make(map[string]interface{})
// 			data["modified_by"] = modifiedBy
// 			if article_name != "" {
// 				data["article_name"] = article_name
// 			}
// 			if state != -1 {
// 				data["state"] = state
// 			}

// 			models.EditArticle(id, data)
// 		} else {
// 			code = e.ERROR_NOT_EXIST_ARTICLE
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
