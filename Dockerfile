FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/weatherBot
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/bot

FROM scratch
ENV botToken=***
ENV apiToken=***
COPY --from=builder /go/bin/bot /go/bin/bot
CMD ["/go/bin/bot"]