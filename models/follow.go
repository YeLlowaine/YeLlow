package models

//Favorite ...
type Follow struct {
	ID          int    `gorm:"primary_key" json:"id"`
	DoctorName  string `json:"doctor_id"`
	PatientName string `json:"patient_id"`
}

//ExistCommentByID ...
func ExistFollowByID(doctor_name, patient_name string) int {
	var follow Follow
	db.Select("id").Where("doctor_name = ? AND patient_name = ?", doctor_name, patient_name).Find(&follow)
	if follow.ID > 0 {
		return follow.ID
	}

	return -1
}

//GetArticleID ...
func GetDoctorID(patient_name string) (follow []Follow) {
	db.Where("patient_name = ?", patient_name).Find(&follow)

	return
}

func GetPatientID(doctor_name string) (follow []Follow) {
	db.Where("doctor_name = ?", doctor_name).Find(&follow)

	return
}

//AddComment ...
func AddFollow(data map[string]string) bool {
	db.Create(&Follow{
		DoctorName:  data["doctor_name"],
		PatientName: data["patient_name"],
	})

	return true
}

//DeleteComment ...
func DeleteFollow(id int) bool {
	db.Where("id = ?", id).Delete(Follow{})

	return true
}
