FROM golang:1.15-alpine AS builder

ARG TAGGED

LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"
LABEL tagged=$TAGGED

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

# RUN go build -o application
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/application

FROM alpine

COPY --from=builder /app/ /app/

WORKDIR /app

CMD ["/app/application"]

