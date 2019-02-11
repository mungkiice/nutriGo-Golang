# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)

# Stage 1 (to create a "build" image, ~850MB)
FROM golang:1.11-alpine3.8 AS builder

COPY . /go/src/github.com/mungkiice/goNutri/
WORKDIR /go/src/github.com/mungkiice/goNutri/
RUN apk update && apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main.sh

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
# Author : Muhammad Iqbal Kurniawan
MAINTAINER Muhammad Iqbal Kurniawan  <m.kurniawanibal@gmail.com>
WORKDIR /root/
COPY --from=builder /go/src/github.com/mungkiice/goNutri/web ./web
COPY --from=builder /go/src/github.com/mungkiice/goNutri/config.json .
COPY --from=builder /go/src/github.com/mungkiice/goNutri/main.sh .

EXPOSE 8000
ENTRYPOINT ["./main.sh"]