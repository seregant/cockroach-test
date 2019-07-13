package main

import (
	gin "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/seregant/cockroach-test/config"
	"github.com/seregant/cockroach-test/controllers"
	"github.com/seregant/cockroach-test/database"
	"github.com/seregant/cockroach-test/middleware"
)

var conf = config.SetConfig()

func main() {
	database.DbInit()
	jabatanController := new(controllers.Jabatan)
	pegawaiController := new(controllers.Pegawai)
	divisiController := new(controllers.Divisi)
	pekerjaanController := new(controllers.Pekerjaan)
	//mid := new(middleware.Default)

	router := gin.Default()
	// corsConfig := middleware.CORSMiddleware()
	router.Use(middleware.CORSMiddleware())

	//route api
	api := router.Group("/api")
	{
		//CRUD jabatan routes
		jabatan := api.Group("/jabatan")
		{
			jabatan.GET("/", jabatanController.GetAllJabatan)
			jabatan.POST("/tambah", jabatanController.TambahJabatan)
			jabatan.POST("/update/:id_jabatan", jabatanController.UpdateJabatan)
			jabatan.GET("/update/:id_jabatan", jabatanController.UpdateJabatan)
			jabatan.POST("/hapus/:id_jabatan", jabatanController.HapusJabatan)
		}

		//CRUD pegawai routes
		pegawai := api.Group("/pegawai")
		{
			pegawai.GET("/", pegawaiController.GetAllPegawai)
			pegawai.POST("/tambah", pegawaiController.TambahPegawai)
			pegawai.GET("/update/:id_pegawai", pegawaiController.UpdatePegawai)
			pegawai.POST("/update/:id_pegawai", pegawaiController.UpdatePegawai)
			pegawai.POST("/hapus/:id_pegawai", pegawaiController.DeletePegawai)
		}

		divisi := api.Group("/divisi")
		{
			divisi.GET("/", divisiController.GetAllDivisi)
			divisi.POST("/tambah", divisiController.TambahDivisi)
			divisi.GET("/update/:id_divisi", divisiController.UpdateDivisi)
			divisi.POST("/update/:id_divisi", divisiController.UpdateDivisi)
			divisi.POST("/hapus/:id_divisi", divisiController.HapusDivisi)
		}

		pekerjaan := api.Group("/pekerjaan")
		{
			pekerjaan.GET("/", pekerjaanController.GetAllPekerjaan)
			pekerjaan.POST("/tambah", pekerjaanController.TambahPekerjaan)
			pekerjaan.POST("/update/:id_pekerjaan", pekerjaanController.UpdatePekerjaan)
			pekerjaan.GET("/update/:id_pekerjaan", pekerjaanController.UpdatePekerjaan)
			pekerjaan.POST("/hapus/:id_pekerjaan", pekerjaanController.DeletePekerjaan)
		}
	}
	router.Run(":" + conf.HttpPort)
}
