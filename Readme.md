## Simple crud restAPI service with go and cockroachDB
How to run :

1. Clone it to your $GOPATH and install govendor

        go get github.com/seregant/cockroach-test
        go get -u github.com/kardianos/govendor

2. Get the libraries
        
        govendor sync

3. Configure your db connection and http port with config/config.go from config.sample file
4. To start the service execute

        go run main.go
