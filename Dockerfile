FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/weatherBot
COPY . .
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates && apk add build-base && apk add sqlite && apk add socat
RUN go get -d -v
RUN  CGO_ENABLED=1 GOOS=linux go build -o /go/bin/bot -a -ldflags '-linkmode external -extldflags "-static"'



FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV botToken=***
ENV apiToken=***
ENV geoToken=***
COPY --from=builder /go/bin/bot /go/bin/bot
CMD ["/go/bin/bot"]
