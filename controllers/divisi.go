package controllers

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/seregant/cockroach-test/database"
	"github.com/seregant/cockroach-test/hash"
	"github.com/seregant/cockroach-test/structs"
)

type Divisi struct{}

func (w *Divisi) GetAllDivisi(c *gin.Context) {
	var arr_divisi []structs.ResDivisi
	var divisi []structs.Divisi

	var db = database.DbConnect()
	defer db.Close()

	db.Find(&divisi)

	for _, data := range divisi {
		arr_divisi = append(arr_divisi, structs.ResDivisi{IDDivisi: data.IDDivisi, NamaDivisi: data.NamaDivisi})
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "success",
		"data":    arr_divisi,
	})
}

func (w *Divisi) TambahDivisi(c *gin.Context) {
	var dataDivisi, _ = c.GetPostForm("NamaDivisi")
	var idDivisi = hash.GenerateIDData()

	db := database.DbConnect()
	defer db.Close()

	add := db.Create(&structs.Divisi{IDDivisi: idDivisi, NamaDivisi: dataDivisi})

	isError(add, c)
}

func (w *Divisi) UpdateDivisi(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.JSON(405, gin.H{
			"status":  "405",
			"message": "method not allowed",
		})
	} else {
		var IDtoEdit = c.Param("id_divisi")
		fmt.Println(IDtoEdit)
		var UpdatedData, _ = c.GetPostForm("NamaDivisi")
		fmt.Println(UpdatedData)
		db := database.DbConnect()
		db.LogMode(true)
		defer db.Close()

		if err := db.Where("divisi_id = ?", IDtoEdit).First(&structs.Divisi{}).Error; err != nil {
			// error handling...
			if gorm.IsRecordNotFoundError(err) {
				fmt.Println(err)
			}
		}

		update := db.Model(&structs.Divisi{}).Where("divisi_id = ?", IDtoEdit).Update("divisi_nama", UpdatedData)

		isError(update, c)
	}
}

func (w *Divisi) HapusDivisi(c *gin.Context) {
	var IdToDel = c.Param("id_divisi")

	var db = database.DbConnect()
	defer db.Close()

	delete := db.Where("divisi_id = ? ", IdToDel).Delete(&structs.Divisi{})

	isError(delete, c)
}
