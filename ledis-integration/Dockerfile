FROM golang:1.9-alpine

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH  

RUN mkdir -p $GOPATH/src/go_exercises/ledis-integration

ADD . $GOPATH/src/go_exercises/ledis-integration

WORKDIR $GOPATH/src/go_exercises/ledis-integration
RUN go build -o ledis-integration . 

ENTRYPOINT $GOPATH/src/go_exercises/ledis-integration

EXPOSE 8080