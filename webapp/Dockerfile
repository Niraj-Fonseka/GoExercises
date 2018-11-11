FROM golang:1.9-alpine

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH  

RUN mkdir -p $GOPATH/src/app 

ADD . $GOPATH/src/app

WORKDIR $GOPATH/src/app 
RUN go build -o webapp . 

ENTRYPOINT $GOPATH/src/app/webapp

EXPOSE 8080