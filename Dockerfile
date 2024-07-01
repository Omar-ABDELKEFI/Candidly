FROM golang:1.16.5

WORKDIR /go/src/tekab-test/

COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/pilu/fresh
RUN go get  github.com/swaggo/swag/cmd/swag@v1.7.0
RUN go get  github.com/arsmn/fiber-swagger/v2
RUN go get gorm.io/datatypes@v1.0.2


ENTRYPOINT swag init --parseDependency --parseInternal && fresh

