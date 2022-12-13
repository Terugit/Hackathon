FROM golang:1.18 as build

WORKDIR /go/src

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main main.go

#EXPOSE 8080

CMD ["/go/src/main"]