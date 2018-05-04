FROM golang:latest

WORKDIR $GOPATH/src/chat
ADD . $GOPATH/src/chat

RUN go get github.com/gorilla/mux && go get github.com/gorilla/securecookie && go get github.com/gorilla/websocket
RUN git clone https://github.com/golang/crypto $GOPATH/src/golang.org/x/crypto
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./chat"]

