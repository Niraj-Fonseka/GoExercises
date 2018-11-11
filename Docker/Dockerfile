FROM golang

ADD . $GOPATH/src/GoExercises/Docker

WORKDIR $GOPATH/src/GoExercises/Docker

RUN cd  $GOPATH/src/GoExercises/Docker; go build -o app
ENTRYPOINT [ "./app" ]