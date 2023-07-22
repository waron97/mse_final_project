FROM golang:1.20



WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./app

CMD ["./app"]