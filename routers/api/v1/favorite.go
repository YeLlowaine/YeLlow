package v1

import (
	"fmt"
	"net/http"

	"github.com/YeLlowaine/YeLlow/models"
	"github.com/YeLlowaine/YeLlow/pkg/e"
	"github.com/YeLlowaine/YeLlow/pkg/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//GetFavoriteArticle 获取收藏文章
func GetFavoriteArticle(c *gin.Context) {
	user_id := com.StrTo(c.Query("user_id")).MustInt()

	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Min(user_id, 1, "id").Message("用户ID必须大于0")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.GetArticleID(user_id)
		for index, val := range ret {
			str := fmt.Sprintf("%d", index)
			data[str] = models.GetArticle(val.ArticleId)
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
	user_id := c.Query("user_id")
	article_id := c.Query("article_id")
	valid := validation.Validation{}

	valid.Required(user_id, "user_id").Message("用户不能为空")
	valid.Required(article_id, "article_id").Message("帖子不能为空")

	maps := make(map[string]int)
	maps["user_id"] = com.StrTo(user_id).MustInt()
	maps["article_id"] = com.StrTo(article_id).MustInt()
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if ret := models.ExistRelationByID(com.StrTo(user_id).MustInt(), com.StrTo(article_id).MustInt()); ret == -1 {
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
	user_id := com.StrTo(c.Query("user_id")).MustInt()
	article_id := com.StrTo(c.Query("article_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(user_id, 1, "id").Message("用户ID必须大于0")
	valid.Min(article_id, 1, "id").Message("帖子ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		if id := models.ExistRelationByID(user_id, article_id); id >= 0 {
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
