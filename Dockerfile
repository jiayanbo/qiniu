FROM golang

RUN mkdir -p /qiniu
COPY . /qiniu
WORKDIR /qiniu
ENTRYPOINT ["./start.sh"]