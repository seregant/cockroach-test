package controllers

import (
	gin "github.com/gin-gonic/gin"
	"github.com/seregant/cockroach-test/database"
	"github.com/seregant/cockroach-test/structs"
)

type Jabatan struct{}

func (w *Jabatan) GetAllJabatan(c *gin.Context) {
	var arr_jabatan []structs.ResJabatan
	var response structs.DataResponse
	var jabatan []structs.Jabatan

	var db = database.DbConnect()
	defer db.Close()

	db.Find(&jabatan)

	for _, jabat := range jabatan {
		arr_jabatan = append(arr_jabatan, structs.ResJabatan{IDJabatan: jabat.IDJabatan, NamaJabatan: jabat.NamaJabatan})
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arr_jabatan

	c.JSON(200, response)
}

func (w *Jabatan) UpdateJabatan(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.JSON(405, gin.H{
			"status":  "405",
			"message": "method not allowed",
		})
	} else {
		var response structs.StdrResponse
		c.Request.ParseForm()

		var IDtoEdit = c.Param("id_jabatan")
		var UpdatedData, _ = c.GetPostForm("NamaJabatan")

		db := database.DbConnect()
		db.LogMode(true)
		defer db.Close()

		update := db.Model(&structs.Jabatan{}).Where("jabatan_id = ?", IDtoEdit).Update("NamaJabatan", UpdatedData)

		if update.GetErrors() != nil {

			type errMsg struct {
				message string
			}

			var errData []errMsg

			for _, err := range update.GetErrors() {
				errData = append(errData, errMsg{message: err.Error()})
			}

			response.Status = 204
			response.Message = errData[0].message

			c.JSON(200, response)
		} else {
			response.Status = 200
			response.Message = "Success"

			c.JSON(200, response)
		}
	}
}
