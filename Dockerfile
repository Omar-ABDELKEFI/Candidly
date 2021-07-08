FROM golang:1.16.5

WORKDIR C:/users/omara/go/src/tekab-test

COPY . .
COPY go.mod .
COPY go.sum .
RUN go get -d -v ./...
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

EXPOSE 8081