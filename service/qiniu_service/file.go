package qiniu_service

import (
	"blogx_server/global"
	file2 "blogx_server/utils/file"
	"blogx_server/utils/hash"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"

	"io"
)

func SendFile(file string) (url string, err error) {
	mac := credentials.NewCredentials(global.Conifg.QiNiu.AccessKey, global.Conifg.QiNiu.SecretKey)
	hashString, err := hash.FileMd5(file)
	if err != nil {
		return
	}
	suffix, _ := file2.ImageSuffixJudge(file)
	fileName := fmt.Sprintf("%s.%s", hashString, suffix)
	key := fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadFile(
		context.Background(),
		file,
		&uploader.ObjectOptions{
			BucketName: global.Conifg.QiNiu.Bucket,
			ObjectName: &key,
			FileName:   fileName,
		}, nil)
	fmt.Println(global.Conifg.QiNiu.Uri)
	return fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Uri, key), nil
}

func sendReader(reader io.Reader) (url string, err error) {
	mac := credentials.NewCredentials(global.Conifg.QiNiu.AccessKey, global.Conifg.QiNiu.SecretKey)
	uid := uuid.New().String()

	fileName := fmt.Sprintf("%s.png", uid)
	key := fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadReader(
		context.Background(),
		reader,
		&uploader.ObjectOptions{
			BucketName: global.Conifg.QiNiu.Bucket,
			ObjectName: &key,
			FileName:   fileName,
		}, nil)
	fmt.Println(global.Conifg.QiNiu.Uri)
	return fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Uri, key), nil
}
