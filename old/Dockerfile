FROM golang:1.7.3-alpine

ENV SOURCES /go/src/github.com/miloofcroton/simple-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT simple-go
