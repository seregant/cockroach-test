package structs

type ResJabatan struct {
	IDJabatan   int    `form:"id" json:"id"`
	NamaJabatan string `form:"nama" json:"nama"`
}

type ResPegawai struct {
	IDPegawai int    `form:"id" json:"id"`
	Nama      string `form:"nama" json:"nama"`
	Alamat    string `form:"alamat" json:"alamat"`
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	Email     string `form:"email" json:"email"`
	JabatanID int    `form:"jabatan_id" json:"jabatan_id"`
	DivisiID  int    `form:"divisi_id" json:"divisi_id"`
}
