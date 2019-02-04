#first stage - builder
FROM golang:1.11.0-stretch as builder

COPY . $GOPATH/src/LoadLocationTest
WORKDIR $GOPATH/src/LoadLocationTest

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o loadLocation 



#second stage 
FROM alpine:latest

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /root/

COPY --from=builder /go/src/LoadLocationTest .

EXPOSE 8080

CMD ["./loadLocation"]
