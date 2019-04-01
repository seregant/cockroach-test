package structs

var prefix = "krywn_"

type Pegawai struct {
	IDPegawai int    `gorm:"primary_key:yes;column:pegawai_id;auto_increment:yes"`
	Nama      string `gorm:"column:pegawai_nama"`
	Alamat    string `gorm:"column:pegawai_alamat"`
}

func (Pegawai) TableName() string {
	return prefix + "pegawai"
}

type Divisi struct {
	IDDivisi   int    `gorm:"primary_key:yes;column:divisi_id;auto_increment:yes"`
	NamaDivisi string `gorm:"column:divisi_nama"`
}

func (Divisi) TableName() string {
	return prefix + "divisi"
}

type Jabatan struct {
	IDJabatan   int    `gorm:"primary_key:yes;column:jabatan_id;auto_increment:yes"`
	NamaJabatan string `gorm:"column:jabatan_nama"`
}

func (Jabatan) TableName() string {
	return prefix + "jabatan"
}
