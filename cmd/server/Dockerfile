FROM golang:1.11-stretch

WORKDIR /go/src/github.com/aren55555/aren55555/cmd/server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["server"]
