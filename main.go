package main

import (
	"github.com/seregant/cockroach-test/middleware"

	gin "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/seregant/cockroach-test/config"
	"github.com/seregant/cockroach-test/controllers"
)

var conf = config.SetConfig()

func main() {
	//database.DbInit()
	jabatanController := new(controllers.Jabatan)
	pegawaiController := new(controllers.Pegawai)

	router := gin.Default()
	router.Use(middleware.ServiceAuth())

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
	}
	router.Run(":" + conf.HttpPort)
}
