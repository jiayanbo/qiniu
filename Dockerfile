FROM golang
RUN mkdir -p /qiniu
COPY . /qiniu
WORKDIR /qiniu
#ENV 设置环境变量
ENV GOPROXY https://goproxy.io 
RUN go build main.go
#CMD ["/qiniu/main", "-k", "test.txt", "-s", "./test.txt", "-b", "sonkwo-gamefile"]
#CMD ["/qiniu/start.sh"]
ENTRYPOINT ["./start.sh"]
