package structs

import "time"

var prefix = "krywn_"

type Pegawai struct {
	IDPegawai string `gorm:"primary_key:yes;column:pegawai_id"`
	Nama      string `gorm:"column:pegawai_nama"`
	Alamat    string `gorm:"column:pegawai_alamat"`
	Username  string `gorm:"column:pegawai_username"`
	Password  string `gorm:"column:pegawai_password"`
	Email     string `gorm:"column:pegawai_email"`
	JabatanID int    `gorm:"column:jabatan_id"`
	DivisiID  int    `gorm:"column:divisi_id"`
}

func (Pegawai) TableName() string {
	return prefix + "pegawai"
}

type Divisi struct {
	IDDivisi   string `gorm:"column:divisi_id"`
	NamaDivisi string `gorm:"column:divisi_nama"`
}

func (Divisi) TableName() string {
	return prefix + "divisi"
}

type Jabatan struct {
	IDJabatan   string `gorm:"column:jabatan_id"`
	NamaJabatan string `gorm:"column:jabatan_nama"`
}

func (Jabatan) TableName() string {
	return prefix + "jabatan"
}

type Pekerjaan struct {
	IDPekerjaan   int       `gorm:"primary_key:yes;column:pekerjaan_id;auto_increment:yes"`
	NamaPekerjaan string    `gorm:"column:pekerjaan_nama"`
	IDPj          int       `gorm:"column:pegawai_id"`
	TimID         time.Time `gorm:"column:team_id"`
	Deadline      string    `gorm:"column:deadline"`
}

func (Pekerjaan) TableName() string {
	return prefix + "pekerjaan"
}

type Team struct {
	IDTeam    time.Time `gorm:"column:team_id"`
	IDPegawai string    `gotm:"column:pegawai_id"`
}

func (Team) TableName() string {
	return prefix + "team"
}
