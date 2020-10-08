FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/abracax/JinTalkBackend
COPY . $GOPATH/src/github.com/abracax/JinTalkBackend
RUN go build .

EXPOSE 2333
ENTRYPOINT ["./JinTalkBackend"]