FROM golang:latest





ENV GOPROXY https://goproxy.cn,direct






RUN go build .





EXPOSE 8000




ENTRYPOINT ["./apisvr"]
