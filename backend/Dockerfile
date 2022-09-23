FROM golang:1.17-bullseye

WORKDIR /backend
ENV GO111MODULE on
ENV TZ "Asia/Tokyo"

RUN apt update; apt -y upgrade

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . ./backend

RUN go install github.com/rubenv/sql-migrate/...@latest

RUN go get -d github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 5000
CMD ["air", "-c", ".air.toml"]
