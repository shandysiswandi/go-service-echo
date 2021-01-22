FROM golang:1.15-alpine AS builder

ARG builder

LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"
LABEL buildername=$builder

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

