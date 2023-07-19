FROM golang:1.20

# RUN apt-get update && apt-get upgrade -y
# RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
# RUN apt-get install -y nodejs 
# RUN npm install -g npm@latest
# RUN npm install -g nodemon 

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./app ./cmd/indexer

CMD ["./app"]