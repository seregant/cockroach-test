## Simple crud restAPI service with go and cockroachDB
Cara menjalankan :

1. Clone program ke $GOPATH dan install govendor

        go get github.com/seregant/cockroach-test
        go get -u github.com/kardianos/govendor

2. Download library dengan govendor
        
        cd ${GOPATH}/src/github.com/seregant/cockroach-test/
        govendor sync

3. Sesuaikan config database pada file `config/config.go`
4. Jalankan service dengan command

        go run main.go

## Menggunakan docker-compose

1. Install docker dan docker-compose pada server
2.  Download file docker-compose.yml atau clone program lalu jalankan command :

        cd ${GOPATH}/src/github.com/seregant/cockroach-test/
        docker-compose -d up

## API Route list

  ##CRUD jabatan :

	GET all data    = /jabatan/
	POST add        = /jabatan/tambah
	POST update     = /jabatan/update/:id_jabatan
	POST Delete     = /jabatan/hapus/:id_jabatan
	
  ##CRUD pegawai :

	GET all data    = /pegawai/
	POST add        = /pegawai/tambah
        POST update     = /pegawai/update/:id_pegawai
	POST Delete     = /pegawai/hapus/:id_pegawai
 
 ##CRUD divisi 
	
	GET all data    = /divisi/
	POST add        = /divisi/tambah
	POST update     = /divisi/update/:id_divisi
	POST Delete     = /divisi/hapus/:id_divisi

 ##CRUD pekerjaan :

	GET all data    = /pekerjaan/
	POST add        = /pekerjaan/tambah
	POST update     = /pekerjaan/update/:id_pekerjaan
	POST Delete     = /pekerjaan/hapus/:id_pekerjaan

 ##POSTMAN collection link :

        https://www.getpostman.com/collections/6eecd6aa31ad2ec3184e