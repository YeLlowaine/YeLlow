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
	Duration         string `json:"duration"`
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

func CheckQuestion(username, question string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, SecurityQuestion: question}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

// AddAuth
func AddAuth(username, password, face_picture, icon, address, security_question, duration string, user_type int) bool {
	// var auth Auth
	db.Create(&Auth{
		Username:         username,
		Password:         password,
		FacePicture:      face_picture,
		Icon:             icon,
		Address:          address,
		SecurityQuestion: security_question,
		Duration:         duration,
		UserType:         user_type,
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

func GetUser(username string) (auth Auth) {
	db.Where("username = ?", username).Find(&auth)

	return
}

func SearchUser(keyword string) (auth []Auth) {
	db.Where("username LIKE ? AND user_type = 1", keyword).Find(&auth)

	return
}

// GetInfo ...
func GetInfo(id int) (auth Auth) {
	db.Where("id = ?", id).First(&auth)

	return
}

// GetInfoFromName ...
func GetInfoFromName(username string) (auth Auth) {
	db.Where("username = ?", username).First(&auth)

	return
}

// GetIcon ...
func GetIcon(username string) (auth Auth) {
	db.Select("icon").Where("username = ?", username).First(&auth)

	return
}

func GetFaceToken(username string) (auth Auth) {
	db.Select("face_picture").Where("username = ?", username).First(&auth)

	return
}

func GetAll() (auth []Auth) {
	db.Select("*").Where("user_type = 1").Find(&auth)

	return
}
