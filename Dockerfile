FROM alpine
LABEL author="Alberto Bregliano"

WORKDIR /app

RUN apk update && apk add --no-cache go && apk add --no-cache git

# RUN git clone https://github.com/albertobregliano/webservertemporizzato.git
# RUN echo $PATH
RUN go install github.com/albertobregliano/webservertemporizzato@v0.2.1

ENV TIMEOUT "20s"
ENV PORT "8080"

#ENTRYPOINT [ "/root/go/bin/webservertemporizzato" ]
CMD /root/go/bin/webservertemporizzato -t $TIMEOUT -p $PORT