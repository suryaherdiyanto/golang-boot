FROM golang:1.23.8-alpine

RUN mkdir /.cache && chmod 777 /.cache

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go install -tags 'postgres,mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["air", "-c", ".air.toml"]
