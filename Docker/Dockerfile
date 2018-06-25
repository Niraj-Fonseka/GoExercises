FROM golang

ADD . /go/src/GoExercises/Docker

RUN cd /go/src/GoExercises/Docker; go build -o dockerApp

ENTRYPOINT [ "./dockerApp" ]