FROM alpine

RUN apk update && apk add go && apk add git
ENV GO111MODULE="on"
ENV GOOS="linux"
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go install github.com/albertobregliano/webservertemporizzato@latest

ENTRYPOINT [ "webservertemporizzato" ]