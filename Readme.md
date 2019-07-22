## Simple crud restAPI service with go and cockroachDB
Cara menjalankan :

1. Clone program ke $GOPATH dan install govendor

        go get github.com/seregant/cockroach-test
        go get -u github.com/kardianos/govendor

2. Download library dengan govendor
        
        govendor sync

3. Sesuaikan config database pada file `config/config.go`
4. Jalankan service dengan command

        go run main.go

## Menggunakan docker-compose

1. Install docker dan docker-compose pada server
2.  Download file docker-compose.yml atau clone program lalu jalankan command :

        docker-compose -d up

## API Route list


