package models

//Auth ...
type Auth struct {
	ID               int    `gorm:"primary_key" json:"id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	UserType         int    `json:"user_type"`
	FacePicture      string `json:"face_picture"`
	Icon             string `json:"icon"`
	Address          string `json:"address"`
	SecurityQuestion string `json:"security_question"`
}

//CheckAuth ...
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

// AddAuth
func AddAuth(username, password, face_picture, icon, address, security_question string, user_type int) bool {
	// var auth Auth
	db.Create(&Auth{
		Username:         username,
		Password:         password,
		UserType:         user_type,
		FacePicture:      face_picture,
		Icon:             icon,
		Address:          address,
		SecurityQuestion: security_question,
	})

	return true
}

//CheckAnswer
func CheckAnswer(username string, security_question string) bool {
	var auth Auth
	//	db.Select("security_question").Where("security_question = ?", security_question).Where("security_question = ?", security_question).First(&auth)
	db.Where("username = ? AND security_question = ?", username, security_question).Find(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

//UpdateAuth ...
func UpdateAuth(username string, data interface{}) bool {
	db.Model(&Auth{}).Where("username = ?", username).Updates(data)

	return true
}

func GetUser(user_id int) (auth []Auth) {
	db.Where("id = ?", user_id).Find(&auth)

	return
}

func SearchUser(keyword string) (auth []Auth) {
	db.Where("username LIKE ? AND user_type = 1", keyword).Find(&auth)

	return
}
