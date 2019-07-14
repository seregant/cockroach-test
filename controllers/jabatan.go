package controllers

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/seregant/cockroach-test/database"
	"github.com/seregant/cockroach-test/hash"
	"github.com/seregant/cockroach-test/structs"
)

type Jabatan struct{}

func (w *Jabatan) GetAllJabatan(c *gin.Context) {
	var arr_jabatan []structs.ResJabatan
	var jabatan []structs.Jabatan

	var db = database.DbConnect()
	defer db.Close()

	db.Find(&jabatan)

	for _, jabat := range jabatan {
		arr_jabatan = append(arr_jabatan, structs.ResJabatan{IDJabatan: jabat.IDJabatan, NamaJabatan: jabat.NamaJabatan})
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "success",
		"data":    arr_jabatan,
	})
}

func (w *Jabatan) TambahJabatan(c *gin.Context) {
	var dataJabatan, _ = c.GetPostForm("NamaJabatan")
	var idJabatan = hash.GenerateIDData()

	db := database.DbConnect()
	defer db.Close()

	add := db.Create(&structs.Jabatan{IDJabatan: "JB" + idJabatan, NamaJabatan: dataJabatan})

	isError(add, c)
}

func (w *Jabatan) UpdateJabatan(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.JSON(405, gin.H{
			"status":  "405",
			"message": "method not allowed",
		})
	} else {
		var IDtoEdit = c.Param("id_jabatan")
		var UpdatedData, _ = c.GetPostForm("NamaJabatan")

		db := database.DbConnect()
		db.LogMode(true)
		defer db.Close()

		update := db.Model(&structs.Jabatan{}).Where("jabatan_id = ? ", IDtoEdit).Update("NamaJabatan", UpdatedData)

		isError(update, c)
	}
}

func (w *Jabatan) HapusJabatan(c *gin.Context) {
	var IdToDel = c.Param("id_jabatan")

	var db = database.DbConnect()
	defer db.Close()

	delete := db.Where("jabatan_id = ? ", IdToDel).Delete(&structs.Jabatan{})

	isError(delete, c)
}

func isError(a *gorm.DB, c *gin.Context) {
	// fmt.Println(a.Value)
	if a.GetErrors() != nil {

		type errMsg struct {
			message string
		}

		var errData []errMsg
		fmt.Println(a.GetErrors())
		for _, err := range a.GetErrors() {
			errData = append(errData, errMsg{message: err.Error()})
		}
		c.JSON(200, gin.H{
			"status":   "200",
			"message":  "operation failed",
			"err_data": errData,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "200",
			"message": "OK",
		})
	}
}
