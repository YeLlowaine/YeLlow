package models

//Favorite ...
type Favorite struct {
	ID        int `gorm:"primary_key" json:"id"`
	UserId    int `json:"user_id"`
	ArticleId int `json:"article_id"`
}

//ExistRelationByID是否已经被收藏 ...
func ExistRelationByID(user_id, article_id int) int {
	var favorite Favorite
	db.Select("id").Where("user_id = ? AND article_id = ?", user_id, article_id).Find(&favorite)
	if favorite.ID > 0 {
		return favorite.ID
	}

	return -1
}

//GetArticleID ...
func GetArticleID(user_id int) (favorite []Favorite) {
	db.Where("user_id = ?", user_id).Find(&favorite)

	return
}

//AddFavorite添加到收藏 ...
func AddFavorite(data map[string]int) bool {
	db.Create(&Favorite{
		UserId:    data["user_id"],
		ArticleId: data["article_id"],
	})

	return true
}

//DeleteFavorite从收藏中删除 ...
func DeleteFavorite(id int) bool {
	db.Where("id = ?", id).Delete(Favorite{})

	return true
}
