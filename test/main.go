package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	key =flag.String("k", "test.txt", "file name")
	localFile =flag.String("s", "./test.txt", "source file path")
	bucket =flag.String("b", "sonkwo-gamefile", "bucket")
)

var accessKey string = "42W5KgVowhDP_NL_j85uvpTARFhsbUwZzfGGS2R5"
var	secretKey string  = "Px04ihNML4IyGNtezVZn0-Q2L4W-dcNXd5Rx8FV_"
func main(){
	flag.Parse()
	uploadFile()
	downloadFile()
}


func uploadFile(){
	//localFile = "E:/work/qiniu/test_upload.txt"
	//bucket = "sonkwo-gamefile"
	//key = "test_upload.txt"
	putPolicy := storage.PutPolicy{
		Scope:               *bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "sonkwo test",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, *key, *localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key,ret.Hash)
}
//http://ohpgpdoop.bkt.clouddn.com/3SwitcheD.mf
func downloadFile(){
	mac := qbox.NewMac(accessKey, secretKey)
	domain := "http://ohpgpdoop.bkt.clouddn.com"
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, *key, deadline)
	fmt.Println(privateAccessURL)

	download( fmt.Sprintf("./%s", *key), privateAccessURL)
}

func download(fileFullPath, url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK{
		log.Println("err code: ", res.StatusCode)
		return
	}
	f, err := os.Create(fileFullPath)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

func compareFile(){


}