FROM golang:1.11
WORKDIR  /go/src/github.com/hirany/web_app/ 
RUN go get -u golang.org/x/vgo
COPY . .
RUN vgo build -o chat
CMD ./chat

