FROM golang:latest
MAINTAINER abracax
ENV GOPROXY https://goproxy.cn,direct

RUN go get -v github.com/gin-gonic/gin
RUN go get github.com/Abracax/JinTalkBackend

WORKDIR $GOPATH/src/github.com/abracax/JinTalkBackend
COPY . $GOPATH/src/github.com/abracax/JinTalkBackend
RUN go build .

EXPOSE 2333
ENTRYPOINT ["./JinTalkBackend"]