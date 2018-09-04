FROM golang:1.11-stretch

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/aren55555/aren55555/cmd/client

COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -v -vendor-only

COPY . .

RUN go get -d ./...
RUN go install ./...

CMD ["client"]
