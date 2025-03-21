// testdata/10.七牛云.go
package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"context"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"os"

	"io"

	//"blogx_server/service/qiniu_service"
	"fmt"
)

func SendReader(reader io.Reader) (url string, err error) {
	mac := credentials.NewCredentials(global.Conifg.QiNiu.AccessKey, global.Conifg.QiNiu.SecretKey)

	//uid := uuid.New().String()
	//fileName := fmt.Sprintf("%s.png", uid)
	key := "blogx/519ca0ef5f4e3e864139d7440d595160.jpg"
	//key := fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: global.Conifg.QiNiu.Bucket,
		ObjectName: &key,
		FileName:   "xx.jpg",
	}, nil)
	return fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Uri, key), err
}

func main() {
	flags.Parse() //参数解析
	//fmt.Println(flags.FlagOptions)
	global.Conifg = core.ReadConf()
	core.InitLogrus()
	file, _ := os.Open("uploads/images001/519ca0ef5f4e3e864139d7440d595160.jpg")
	url, err := SendReader(file)
	fmt.Println(url, err)
}
