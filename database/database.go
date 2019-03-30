package database

import (
	"fmt"
	"log"

	"github.com/seregant/cockroach-test/config"
	"github.com/seregant/cockroach-test/structs"

	"github.com/jinzhu/gorm"
)

func DbConnect() *gorm.DB {
	var conf = config.SetConfig()
	var addr = "postgresql://" + conf.User + "@" + conf.Host + ":" + conf.Port + "/" + conf.DbName + "?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DbInit() {

	var db = DbConnect()
	defer db.Close()

	fmt.Println("Creating tables...")

	db.AutoMigrate(&structs.Pegawai{})
	db.AutoMigrate(&structs.Divisi{})
	db.AutoMigrate(&structs.Jabatan{})

	fmt.Println("Initiating tables...")

	db.Create(&structs.Jabatan{IDJabatan: 1, NamaJabatan: "direktur"})
	db.Create(&structs.Jabatan{IDJabatan: 2, NamaJabatan: "manager"})
	db.Create(&structs.Jabatan{IDJabatan: 3, NamaJabatan: "supervisor"})
	db.Create(&structs.Jabatan{IDJabatan: 4, NamaJabatan: "staff"})

}
