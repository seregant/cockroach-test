package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seregant/cockroach-test/database"
	"github.com/seregant/cockroach-test/hash"
	"github.com/seregant/cockroach-test/structs"
)

type Pegawai struct{}

func (w *Pegawai) GetAllPegawai(c *gin.Context) {
	var arr_pegawai []structs.ResPegawai
	var pegawai []structs.Pegawai
	var jabatan structs.Jabatan
	var divisi structs.Divisi

	db := database.DbConnect()
	defer db.Close()

	db.Find(&pegawai)

	for _, data := range pegawai {
		fmt.Println(data.IDPegawai)
		db.Where("jabatan_id = ?", data.JabatanID).Find(&jabatan)
		db.Where("divisi_id = ?", data.DivisiID).Find(&divisi)
		arr_pegawai = append(arr_pegawai, structs.ResPegawai{
			IDPegawai: data.IDPegawai,
			Nama:      data.Nama,
			Alamat:    data.Alamat,
			Username:  data.Username,
			Password:  data.Password,
			Email:     data.Email,
			JabatanID: jabatan.NamaJabatan,
			DivisiID:  divisi.NamaDivisi,
		})
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
	var username, _ = c.GetPostForm("username")
	var password, _ = c.GetPostForm("password")
	var email, _ = c.GetPostForm("email")
	var divisi, _ = c.GetPostForm("divisi")
	var jabatan, _ = c.GetPostForm("jabatan")

	fmt.Println(nama, alamat, username, password, email, divisi, jabatan)

	divisiVal, _ := strconv.Atoi(divisi)
	jabatanVal, _ := strconv.Atoi(jabatan)

	passVal, _ := hash.HashPassword(password)

	db := database.DbConnect()
	defer db.Close()

	add := db.Create(&structs.Pegawai{
		IDPegawai: "PG" + hash.GenerateIDData(),
		Nama:      nama,
		Alamat:    alamat,
		Username:  username,
		Password:  passVal,
		Email:     email,
		JabatanID: jabatanVal,
		DivisiID:  divisiVal,
	})

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
		var username, _ = c.GetPostForm("username")
		var password, _ = c.GetPostForm("password")
		var email, _ = c.GetPostForm("email")
		var divisi, _ = c.GetPostForm("divisi")
		var jabatan, _ = c.GetPostForm("jabatan")

		passVal, _ := hash.HashPassword(password)

		db := database.DbConnect()
		db.LogMode(true)
		defer db.Close()

		update := db.Exec("UPDATE krywn_pegawai SET pegawai_nama='" + dataNama + "', pegawai_alamat='" + dataAlamat + "', pegawai_username='" + username + "', pegawai_password='" + passVal + "', pegawai_email='" + email + "', jabatan_id=" + jabatan + ", divisi_id=" + divisi + " WHERE pegawai_id = " + IDtoEdit)

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
