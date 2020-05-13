package models

//Favorite ...
type Follow struct {
	ID        int `gorm:"primary_key" json:"id"`
	DoctorId  int `json:"doctor_id"`
	PatientId int `json:"patient_id"`
}

//ExistCommentByID ...
func ExistFollowByID(doctor_id, patient_id int) int {
	var follow Follow
	db.Select("id").Where("doctor_id = ? AND patient_id = ?", doctor_id, patient_id).Find(&follow)
	if follow.ID > 0 {
		return follow.ID
	}

	return -1
}

//GetArticleID ...
func GetDoctorID(patient_id int) (follow []Follow) {
	db.Where("patient_id = ?", patient_id).Find(&follow)

	return 
}

func GetPatientID(doctor_id int) (follow []Follow) {
	db.Where("doctor_id = ?", doctor_id).Find(&follow)

	return
}

//AddComment ...
func AddFollow(data map[string]int) bool {
	db.Create(&Follow{
		DoctorId:  data["doctor_id"],
		PatientId: data["patient_id"],
	})

	return true
}

//DeleteComment ...
func DeleteFollow(id int) bool {
	db.Where("id = ?", id).Delete(Follow{})

	return true
}
