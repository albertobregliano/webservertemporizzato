FROM alpine

RUN apk update && apk add go && apk add git

# RUN git clone https://github.com/albertobregliano/webservertemporizzato.git
RUN GOBIN=/usr/bin && go install github.com/albertobregliano/webservertemporizzato@v1

ENTRYPOINT [ "/webservertemporizzato/webservertemporizzato" ]