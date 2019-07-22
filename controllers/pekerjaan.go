package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seregant/cockroach-test/database"
	"github.com/seregant/cockroach-test/structs"
)

type Pekerjaan struct{}

func (w *Pekerjaan) GetAllPekerjaan(c *gin.Context) {
	var arr_pekerjaan []structs.ResPekerjaan
	var pekerjaan []structs.Pekerjaan
	var team []structs.Team
	var teamMember []string

	db := database.DbConnect()
	defer db.Close()

	db.Find(&pekerjaan)

	for _, data := range pekerjaan {
		db.Where("team_id = ?", data.TimID).Find(&team)
		for _, dataTeam := range team {
			var pegawaiStruct structs.Pegawai
			db.Where("pegawai_id = ?", dataTeam.IDPegawai).First(&pegawaiStruct)
			teamMember = append(teamMember, pegawaiStruct.Nama)
		}

		var pegawaiStruct structs.Pegawai
		db.Where("pegawai_id = ?", data.IDPj).First(&pegawaiStruct)

		arr_pekerjaan = append(arr_pekerjaan, structs.ResPekerjaan{
			IDPekerjaan:   data.IDPekerjaan,
			NamaPekerjaan: data.NamaPekerjaan,
			NamaPj:        pegawaiStruct.Nama,
			Tim:           teamMember,
			Deadline:      data.Deadline,
		})
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "success",
		"data":    arr_pekerjaan,
	})

}

func (w *Pekerjaan) TambahPekerjaan(c *gin.Context) {

	var Nama, _ = c.GetPostForm("nama")
	var PjId, _ = c.GetPostForm("pj_id")
	var Deadline, _ = c.GetPostForm("deadline")
	var TeamMemberId = c.PostFormArray("team_member_id")

	PjIdInt, _ := strconv.Atoi(PjId)

	db := database.DbConnect()
	defer db.Close()

	teamId := time.Now()
	teamId.Format("20060102150405")

	db.Create(&structs.Pekerjaan{
		NamaPekerjaan: Nama,
		IDPj:          PjIdInt,
		TimID:         teamId,
		Deadline:      Deadline,
	})

	for _, dataTeamId := range TeamMemberId {
		var pegawaiStruct structs.Pegawai
		db.Where("pegawai_id = ?", dataTeamId).First(&pegawaiStruct)
		fmt.Println(pegawaiStruct.IDPegawai)

		db.Create(&structs.Team{
			IDTeam:    teamId,
			IDPegawai: pegawaiStruct.IDPegawai,
		})
	}
}

func (w *Pekerjaan) UpdatePekerjaan(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.JSON(405, gin.H{
			"status":  "405",
			"message": "method not allowed",
		})
	} else {
		var IDToEdit = c.Param("id_pekerjaan")
		var Nama, _ = c.GetPostForm("nama")
		var PjId, _ = c.GetPostForm("pj_id")
		var Deadline, _ = c.GetPostForm("deadline")
		var TeamMemberId = c.PostFormArray("team_member_id")

		var team structs.Team
		var pekerjaan structs.Pekerjaan

		db := database.DbConnect()
		// db.LogMode(true)
		defer db.Close()

		db.Where("pekerjaan_id = ?", IDToEdit).First(&pekerjaan)
		db.Where("team_id = ?", pekerjaan.TimID).Delete(&team)

		db.Exec("UPDATE krywn_pekerjaan SET pekerjaan_nama='" + Nama + "', pegawai_id='" + PjId + "', deadline='" + Deadline + "' WHERE pekerjaan_id = " + IDToEdit)

		for _, data := range TeamMemberId {
			var pegawaiStruct structs.Pegawai
			db.Where("pegawai_id = ?", data).First(&pegawaiStruct)
			db.Create(&structs.Team{
				IDTeam:    pekerjaan.TimID,
				IDPegawai: pegawaiStruct.IDPegawai,
			})
		}
	}
}

func (w *Pekerjaan) DeletePekerjaan(c *gin.Context) {
	var IdToDel = c.Param("id_pekerjaan")
	var pekerjaan structs.Pekerjaan

	var db = database.DbConnect()
	defer db.Close()
	db.Where("team_id = ?", IdToDel).First(&pekerjaan)

	db.Where("team_id = ?", pekerjaan.TimID).Delete(&structs.Team{})

	db.Where("pekerjaan_id = ? ", IdToDel).Delete(&structs.Pekerjaan{})
}
