FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/weatherBot
COPY . .
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/bot

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV botToken=***
ENV apiToken=***

COPY --from=builder /go/bin/bot /go/bin/bot
CMD ["/go/bin/bot"]