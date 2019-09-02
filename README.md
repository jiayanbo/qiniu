# qiniu
upload、download
docker build -t qiniu:1.0.0 -f ./Dockerfile .
docker run --name qiniu -v /home/sonkwo/qiniu/qiniu/:/qiniu qiniu:1.0.0
