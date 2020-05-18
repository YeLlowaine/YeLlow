package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Article ...
type Article struct {
	Model
	Auth        Auth   `json:"auth"`
	ArticleName string `json:"article_name"`
	CreatedBy   string `json:"created_by"`
	State       int    `json:"state"`
	Content     string `json:"content"`
	DeletedOn   int    `json:"deleted_on"`
}

//GetArticles ...
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

//GetArticlesFromSomeone ...
func GetArticlesFromSomeone(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Auth").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

//GetArticle ...
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)

	return
}

func GetArticleByname(name string) (article Article) {
	db.Where("article_name = ?", name).First(&article)

	return
}

//GetArticleTotal ...
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

//ExistArticleByName ...
func ExistArticleByName(article_name string) bool {
	var article Article
	db.Select("id").Where("article_name = ?", article_name).First(&article)
	if article.ID > 0 {
		return true
	}

	return false
}

//AddArticle ...
func AddArticle(article_name string, state int, createdBy string, content string) bool {
	db.Create(&Article{
		ArticleName: article_name,
		State:       state,
		CreatedBy:   createdBy,
		Content:     content,
	})

	return true
}

//BeforeCreate ...
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ...
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

//ExistArticleByID ...
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}

	return false
}

//DeleteArticle ...
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})

	return true
}

func SearchArticle(keyword string) (article []Article) {
	db.Where("created_by LIKE ? OR article_name LIKE ?", keyword, keyword).Find(&article)

	return
}

// //EditArticle ...
// func EditArticle(id int, data interface{}) bool {
// 	db.Model(&Article{}).Where("id = ?", id).Updates(data)

// 	return true
// }
