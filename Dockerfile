FROM golang
WORKDIR  /go/src/github.com/hirany/web_app/ 
COPY . .
RUN go get github.com/gorilla/websocket
RUN go build -o chat
CMD ./chat

