package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/seregant/cockroach-test/database"
	"github.com/seregant/cockroach-test/structs"
)

type Pegawai struct{}

func (w *Pegawai) GetAllPegawai(c *gin.Context) {
	var arr_pegawai []structs.ResPegawai
	var pegawai []structs.Pegawai

	db := database.DbConnect()
	defer db.Close()

	db.Find(&pegawai)

	for _, data := range pegawai {
		arr_pegawai = append(arr_pegawai, structs.ResPegawai{IDPegawai: data.IDPegawai, Nama: data.Nama, Alamat: data.Alamat})
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "success",
		"data":    arr_pegawai,
	})

}

func (w *Pegawai) TambahPegawai(c *gin.Context) {
	var nama, _ = c.GetPostForm("nama")
	var alamat, _ = c.GetPostForm("alamat")

	db := database.DbConnect()
	defer db.Close()

	add := db.Create(&structs.Pegawai{Nama: nama, Alamat: alamat})

	isError(add, c)

}

func (w *Pegawai) UpdatePegawai(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.JSON(405, gin.H{
			"status":  "405",
			"message": "method not allowed",
		})
	} else {
		var IDtoEdit = c.Param("id_pegawai")
		var dataNama, _ = c.GetPostForm("nama")
		var dataAlamat, _ = c.GetPostForm("alamat")

		db := database.DbConnect()
		db.LogMode(true)
		defer db.Close()

		update := db.Exec("UPDATE krywn_pegawai SET pegawai_nama='" + dataNama + "', pegawai_alamat='" + dataAlamat + "' WHERE pegawai_id = " + IDtoEdit)

		isError(update, c)
	}
}

func (w *Pegawai) DeletePegawai(c *gin.Context) {
	var IdToDel = c.Param("id_pegawai")

	var db = database.DbConnect()
	defer db.Close()

	delete := db.Where("pegawai_id = ? ", IdToDel).Delete(&structs.Pegawai{})

	isError(delete, c)
}
