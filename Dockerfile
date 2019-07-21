FROM golang:rc-alpine
RUN apk add git --no-cache
RUN  go get github.com/seregant/cockroach-test
RUN ls -al $GOPATH/src/github.com/seregant/cockroach-test
WORKDIR  $GOPATH/src/github.com/seregant/cockroach-test
RUN go get -u github.com/kardianos/govendor
RUN govendor sync -v
CMD [ "go","run","main.go" ]
EXPOSE 1234