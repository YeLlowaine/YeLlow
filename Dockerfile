FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/YeLlowaine/YeLlow
COPY . $GOPATH/src/github.com/YeLlowaine/YeLlow
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./YeLlow"]
