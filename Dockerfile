FROM golang

ADD . /go/src/github.com/harrisonchen/golastic
RUN go get github.com/harrisonchen/golastic
RUN go install github.com/harrisonchen/golastic
ENTRYPOINT /go/bin/golastic

EXPOSE 8080
