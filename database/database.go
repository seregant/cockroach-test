package database

import (
	"fmt"
	"log"
	"time"

	"github.com/seregant/cockroach-test/config"
	"github.com/seregant/cockroach-test/hash"
	"github.com/seregant/cockroach-test/structs"

	"github.com/jinzhu/gorm"
)

var conf = config.SetConfig()

func DbConnect() *gorm.DB {
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

	db.Exec("CREATE DATABASE " + conf.DbName)

	fmt.Println("Creating tables...")

	db.AutoMigrate(&structs.Jabatan{})
	db.AutoMigrate(&structs.Pegawai{})
	db.AutoMigrate(&structs.Divisi{})
	db.AutoMigrate(&structs.Pekerjaan{})
	db.AutoMigrate(&structs.Team{})

	fmt.Println("Initiating tables...")

	db.Create(&structs.Jabatan{IDJabatan: "JB" + hash.GenerateIDData(), NamaJabatan: "direktur"})
	time.Sleep(1 * time.Second)
	db.Create(&structs.Jabatan{IDJabatan: "JB" + hash.GenerateIDData(), NamaJabatan: "manager"})
	time.Sleep(1 * time.Second)
	db.Create(&structs.Jabatan{IDJabatan: "JB" + hash.GenerateIDData(), NamaJabatan: "supervisor"})
	time.Sleep(1 * time.Second)
	db.Create(&structs.Jabatan{IDJabatan: "JB" + hash.GenerateIDData(), NamaJabatan: "staff"})

	db.Create(&structs.Divisi{IDDivisi: "DV" + hash.GenerateIDData(), NamaDivisi: "keuangan"})
	time.Sleep(1 * time.Second)
	db.Create(&structs.Divisi{IDDivisi: "DV" + hash.GenerateIDData(), NamaDivisi: "hrd"})
	time.Sleep(1 * time.Second)
	db.Create(&structs.Divisi{IDDivisi: "DV" + hash.GenerateIDData(), NamaDivisi: "sales"})
	time.Sleep(1 * time.Second)
	db.Create(&structs.Divisi{IDDivisi: "DV" + hash.GenerateIDData(), NamaDivisi: "teknis"})
}
