package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

//Comment ...
type Comment struct {
	Model 

	ArticleID int `json:"article_id" gorm:"index"`
	Article   Article `json:"article"`

	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	State      int    `json:"state"`

}

//ExistCommentByID ...
func ExistCommentByID(id int) bool {
	var comment Comment
	db.Select("id").Where("id = ?", id).First(&comment)

	if comment.ID > 0 {
		return true
	}

	return false
}

//GetCommentTotal ...
func GetCommentTotal(maps interface{}) (count int) {
	db.Model(&Comment{}).Where(maps).Count(&count)

	return
}

//GetComments ...
func GetComments(pageNum int, pageSize int, maps interface{}) (comments []Comment) {
	db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&comments)

	return
}

//GetComment ...
func GetComment(id int) (comment Comment) {
	db.Where("id = ?", id).First(&comment)
	db.Model(&comment).Related(&comment.Article)

	return
}


//AddComment ...
func AddComment(data map[string]interface{}) bool {
	db.Create(&Comment{
		ArticleID:     data["article_id"].(int),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

// //DeleteComment ...
// func DeleteComment(id int) bool {
// 	db.Where("id = ?", id).Delete(Comment{})

// 	return true
// }

//BeforeCreate ...
func (comment *Comment) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//BeforeUpdate ...
func (comment *Comment) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
