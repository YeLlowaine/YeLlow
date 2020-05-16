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

//GetFavoriteArtic

// AddFavorite
func AddFollow(c *gin.Context) {
	doctor_id := c.Query("doctor_id")
	patient_id := c.Query("patient_id")
	valid := validation.Validation{}

	valid.Required(doctor_id, "doctor_id").Message("医生不能为空")
	valid.Required(patient_id, "patient_id").Message("病人不能为空")

	maps := make(map[string]int)
	maps["doctor_id"] = com.StrTo(doctor_id).MustInt()
	maps["patient_id"] = com.StrTo(patient_id).MustInt()
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if ret := models.ExistFollowByID(com.StrTo(doctor_id).MustInt(), com.StrTo(patient_id).MustInt()); ret == -1 {
			code = e.SUCCESS
			models.AddFollow(maps)
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

//DeleteFavorite
func DeleteFollow(c *gin.Context) {
	doctor_id := com.StrTo(c.Query("doctor_id")).MustInt()
	patient_id := com.StrTo(c.Query("patient_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(doctor_id, 1, "id").Message("医生ID必须大于0")
	valid.Min(patient_id, 1, "id").Message("病人ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		if id := models.ExistFollowByID(doctor_id, patient_id); id >= 0 {
			logging.Info("id: %s=", id)
			models.DeleteFollow(id)
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

//
func GetDoctor(c *gin.Context) {
	patient_id := com.StrTo(c.Query("patient_id")).MustInt()
	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Min(patient_id, 1, "id").Message("病人ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.GetDoctorID(patient_id)
		for index, val := range ret {
			str := fmt.Sprintf("%d", index)
			data[str] = models.GetUser(val.DoctorId)
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

//
func GetPatient(c *gin.Context) {
	doctor_id := com.StrTo(c.Query("doctor_id")).MustInt()
	data := make(map[string]interface{})
	valid := validation.Validation{}
	valid.Min(doctor_id, 1, "id").Message("病人ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.GetPatientID(doctor_id)
		for index, val := range ret {
			str := fmt.Sprintf("%d", index)
			data[str] = models.GetUser(val.PatientId)
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
