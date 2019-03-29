package structs

var prefix = "krywn_"

type Pegawai struct {
	IDPegawai int    `gorm:"primary_key;column:pegawai_id;AUTO_INCREMENT;size:6"`
	Nama      string `gorm:"column:pegawai_nama;size:25"`
	Alamat    string `gorm:"column:pegawai_alamat;size:150"`
}

func (Pegawai) TableName() string {
	return prefix + "pegawai"
}

type Divisi struct {
	IDDivisi   int    `gorm:"primary_key;column:divisi_id;AUTO_INCREMENT;size:6"`
	NamaDivisi string `gorm:"column:divisi_nama;size:50"`
}

func (Divisi) TableName() string {
	return prefix + "divisi"
}

type Jabatan struct {
	IDJabatan   int    `gorm:"primary_key;column:jabatan_id;AUTO_INCREMENT;size:6"`
	NamaJabatan string `gorm:"column:jabatan_nama;size:50"`
}

func (Jabatan) TableName() string {
	return prefix + "jabatan"
}
