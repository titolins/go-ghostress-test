FROM golang:stretch

WORKDIR /go/src/github.com/titolins/go-ghostress-test/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-ghostress-test"]
