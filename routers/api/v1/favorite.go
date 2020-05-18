package v1

import (
	"fmt"
	"net/http"

	"github.com/YeLlowaine/YeLlow/models"
	"github.com/YeLlowaine/YeLlow/pkg/e"
	"github.com/YeLlowaine/YeLlow/pkg/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//GetFavoriteArticle 获取收藏文章
func GetFavoriteArticle(c *gin.Context) {
	user_name := c.Query("user_name")

	data := make(map[string]interface{})
	valid := validation.Validation{}

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.GetArticleID(user_name)
		for index, val := range ret {
			str := fmt.Sprintf("%d", index)
			data[str] = models.GetArticleByname(val.ArticleName)
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

// AddFavorite
func AddFavorite(c *gin.Context) {
	user_name := c.Query("user_name")
	article_name := c.Query("article_name")
	valid := validation.Validation{}

	maps := make(map[string]string)
	maps["user_name"] = user_name
	maps["article_name"] = article_name
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if ret := models.ExistRelationByID(user_name, article_name); ret == -1 {
			code = e.SUCCESS
			models.AddFavorite(maps)
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

//DeleteFavorite 删除收藏
func DeleteFavorite(c *gin.Context) {
	user_name := c.Query("user_name")
	article_name := c.Query("article_name")

	valid := validation.Validation{}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		if id := models.ExistRelationByID(user_name, article_name); id >= 0 {
			logging.Info("id: %s=", id)
			models.DeleteFavorite(id)
		} else {
			logging.Info("id: %s=", id)
			code = e.ERROR_NOT_EXIST_ARTICLE
			models.DeleteFavorite(id)
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
