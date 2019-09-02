FROM golang
RUN mkdir -p /qiniu/test
COPY . /qiniu/test
WORKDIR /qiniu/test
#ENV 设置环境变量
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on 
RUN go build main.go
#CMD ["/qiniu/main", "-k", "test.txt", "-s", "./test.txt", "-b", "sonkwo-gamefile"]
#CMD ["/qiniu/start.sh"]
ENTRYPOINT ["./start.sh"]
