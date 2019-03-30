## Simple crud restAPI service with go and cockroachDB
How to run :

1. Clone it to your $GOPATH and install govendor

        go get github.com/seregant/cockroach-test
        go get -u github.com/kardianos/govendor

2. Get the libraries
        
        govendor sync

3. Configure your db connection and http port at config/config.go
4. To start the service execute

        go run main.go


### Api Routes

    GET /jabatan                 --> get all data jabatan
    POST /jabatan/update/{id}    --> update a jabatan's row