FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/nexmo-community/nexmo-go

WORKDIR /go/src/github.com/heaptracetechnology/microservice-nexmo

ADD . /go/src/github.com/heaptracetechnology/microservice-nexmo

RUN go install github.com/heaptracetechnology/microservice-nexmo

ENTRYPOINT microservice-nexmo

EXPOSE 3000