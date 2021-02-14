FROM golang:1.15-alpine AS builder

ARG TAGGED

LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"
LABEL tagged=$TAGGED

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

# RUN go build -o application
RUN CGO_ENABLED=0 GOGC=off GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/application
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/application
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/application

FROM alpine

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata

COPY --from=builder /app/application /app/application
COPY --from=builder /app/.env /app/.env

WORKDIR /app

CMD ["/app/application"]

