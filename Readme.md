## Simple crud restAPI service with go and cockroachDB
How to run :

    1. Clone it to your $GOPATH
    2. Get the libraries
        
        go get -u github.com/lib/pq         --> PostgreSQL driver (compatible with cockroachDB)
        go get github.com/gorilla/mux       --> router library
        go get -u github.com/jinzhu/gorm    --> ORM library

    3. To start the service execute

        go run main.go