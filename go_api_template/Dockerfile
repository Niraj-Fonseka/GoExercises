FROM golang:1.9.1

ENV http_proxy=http://thd-svr-proxy-qa.homedepot.com:9090
ENV HTTP_PROXY=http://thd-svr-proxy-qa.homedepot.com:9090

COPY . $GOPATH/src/APP_NAME
WORKDIR $GOPATH/src/APP_NAME/app

RUN go get
RUN go install

EXPOSE 8080

CMD go run main.go

ENV http_proxy=
ENV HTTP_PROXY=
