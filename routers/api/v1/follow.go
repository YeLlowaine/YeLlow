package v1

import (
	"fmt"
	"net/http"

	"github.com/YeLlowaine/YeLlow/models"
	"github.com/YeLlowaine/YeLlow/pkg/e"
	"github.com/YeLlowaine/YeLlow/pkg/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//GetFavoriteArtic

// AddFavorite
func AddFollow(c *gin.Context) {
	doctor_name := c.Query("doctor_name")
	patient_name := c.Query("patient_name")
	valid := validation.Validation{}

	valid.Required(doctor_name, "doctor_name").Message("医生不能为空")
	valid.Required(patient_name, "patient_name").Message("病人不能为空")

	maps := make(map[string]string)
	maps["doctor_name"] = doctor_name
	maps["patient_name"] = patient_name
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if ret := models.ExistFollowByID(doctor_name, patient_name); ret == -1 {
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
	doctor_name := c.Query("doctor_name")
	patient_name := c.Query("patient_name")

	valid := validation.Validation{}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		if id := models.ExistFollowByID(doctor_name, patient_name); id >= 0 {
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
	patient_name := c.Query("patient_name")
	data := make(map[string]interface{})
	valid := validation.Validation{}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.GetDoctorID(patient_name)
		for index, val := range ret {
			str := fmt.Sprintf("%d", index)
			data[str] = models.GetUser(val.DoctorName)
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
	doctor_name := c.Query("doctor_name")
	data := make(map[string]interface{})
	valid := validation.Validation{}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		ret := models.GetPatientID(doctor_name)
		for index, val := range ret {
			str := fmt.Sprintf("%d", index)
			data[str] = models.GetUser(val.PatientName)
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
