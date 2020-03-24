FROM golang:alpine

RUN mkdir -p /go/src/ssh-honeypot

RUN apk update && apk add git

ADD *.go /go/src/ssh-honeypot/
RUN go get ssh-honeypot
RUN go install ssh-honeypot

ENV HONEYPOT_ADDR 0.0.0.0
ENV HONEYPOT_PORT 22
ENV HONEYPOT_LOGTYPE stdout
ENV HONEYPOT_LOGDEST /honeypot/honeypot.log
ENV HONEYPOT_LOGFORMAT string
ENV HONEYPOT_SERVEREY /honeypot/honeypot_rsa

EXPOSE ${HONEYPOT_PORT}

CMD /go/bin/ssh-honeypot
