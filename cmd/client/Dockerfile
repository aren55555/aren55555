FROM golang:1.11-stretch

WORKDIR /go/src/github.com/aren55555/aren55555/cmd/client
COPY . .

RUN go get -d ./...
RUN go install ./...

CMD ["client"]
