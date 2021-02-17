FROM golang:alpine AS builder
ARG TAGGED
LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"
LABEL tagged=$TAGGED
WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOGC=off GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o /build/binary

FROM alpine
LABEL maintainer="SHANDY SISWANDI <shandysiswandi@gmail.com>"
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata
COPY --from=builder /build/binary /app/binary
COPY --from=builder /build/.env /app/.env
WORKDIR /app
ENTRYPOINT ["/app/binary"]
