FROM golang:latest


WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
run go build -v -o /usr/local/bin/app .

CMD ["app"]