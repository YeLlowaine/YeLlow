package models

//Favorite ...
type Favorite struct {
	ID        int `gorm:"primary_key" json:"id"`
	UserName    string `json:"user_id"`
	ArticleName string `json:"article_id"`
}

//ExistRelationByID是否已经被收藏 ...
func ExistRelationByID(user_name, article_name string) int {
	var favorite Favorite
	db.Select("id").Where("user_name = ? AND article_name = ?", user_name, article_name).Find(&favorite)
	if favorite.ID > 0 {
		return favorite.ID
	}

	return -1
}

//GetArticleID ...
func GetArticleID(user_name string) (favorite []Favorite) {
	db.Where("user_name = ?", user_name).Find(&favorite)

	return
}

//AddFavorite添加到收藏 ...
func AddFavorite(data map[string]string) bool {
	db.Create(&Favorite{
		UserName:    data["user_name"],
		ArticleName: data["article_name"],
	})

	return true
}

//DeleteFavorite从收藏中删除 ...
func DeleteFavorite(id int) bool {
	db.Where("id = ?", id).Delete(Favorite{})

	return true
}
