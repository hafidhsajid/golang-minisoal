# FROM golang:1.16-buster AS builder

# WORKDIR /app
# COPY go.* ./
# RUN go mod download
# COPY *.go ./
# RUN go build -o /4
# EXPOSE 8080
# ENTRYPOINT ["/4"]

FROM golang:alpine
WORKDIR /takehometest

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./takehometest