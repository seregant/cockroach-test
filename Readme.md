## Simple crud restAPI service with go and cockroachDB
How to run :

1. Clone it to your $GOPATH
2. Get the libraries
        
        go get -u github.com/lib/pq         --> PostgreSQL driver (compatible with cockroachDB)
        go get github.com/gorilla/mux       --> router library
        go get -u github.com/jinzhu/gorm    --> ORM library

3. Configure your db connection and http port at config/config.go
4. To start the service execute

        go run main.go


### Api Routes

    GET /jabatan                 --> get all data jabatan
    POST /jabatan/update/{id}    --> update a jabatan's row