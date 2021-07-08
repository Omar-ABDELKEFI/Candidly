FROM golang:1.16.5

WORKDIR /go/src/tekab-test/

COPY . .
COPY go.mod .
COPY go.sum .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go mod download
RUN go get github.com/pilu/fresh

ENTRYPOINT fresh

