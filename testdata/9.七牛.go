package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/qiniu_service"
	"fmt"
)

//func SendFile(file string) (url string, err error) {
//	mac := credentials.NewCredentials(global.Conifg.QiNiu.AccessKey, global.Conifg.QiNiu.SecretKey)
//	hashString, err := hash.FileMd5(file)
//	if err != nil {
//		return
//	}
//	suffix, _ := file2.ImageSuffixJudge(file)
//	fileName := fmt.Sprintf("%s.%s", hashString, suffix)
//	key := fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Prefix, fileName)
//	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
//		Options: http_client.Options{
//			Credentials: mac,
//		},
//	})
//	err = uploadManager.UploadFile(
//		context.Background(),
//		file,
//		&uploader.ObjectOptions{
//			BucketName: global.Conifg.QiNiu.Bucket,
//			ObjectName: &key,
//			FileName:   fileName,
//		}, nil)
//	fmt.Println(global.Conifg.QiNiu.Uri)
//	return fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Uri, key), nil
//}
//
//func sendReader(reader io.Reader) (url string, err error) {
//	mac := credentials.NewCredentials(global.Conifg.QiNiu.AccessKey, global.Conifg.QiNiu.SecretKey)
//	uid := uuid.New().String()
//
//	fileName := fmt.Sprintf("%s.png", uid)
//	key := fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Prefix, fileName)
//	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
//		Options: http_client.Options{
//			Credentials: mac,
//		},
//	})
//	err = uploadManager.UploadReader(
//		context.Background(),
//		reader,
//		&uploader.ObjectOptions{
//			BucketName: global.Conifg.QiNiu.Bucket,
//			ObjectName: &key,
//			FileName:   fileName,
//		}, nil)
//	fmt.Println(global.Conifg.QiNiu.Uri)
//	return fmt.Sprintf("%s/%s", global.Conifg.QiNiu.Uri, key), nil
//}
//
//func GenToken() (token string, err error) {
//	mac := credentials.NewCredentials(global.Conifg.QiNiu.AccessKey, global.Conifg.QiNiu.SecretKey)
//	putPolicy, err := uptoken.NewPutPolicy(global.Conifg.QiNiu.Bucket, time.Now().Add(1*time.Minute))
//	if err != nil {
//		return
//	}
//	token, err = uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
//	if err != nil {
//		return
//	}
//	return
//}

func main() {

	flags.Parse() //参数解析
	//fmt.Println(flags.FlagOptions)
	global.Conifg = core.ReadConf()
	core.InitLogrus()

	//url, err := SendFile("uploads/images001/ava.jpg")

	//fmt.Println(url, err)
	//file, _ := os.Open("uploads/images001/ava.jpg")
	//url, err := sendReader(file)
	//fmt.Println(url, err)
	fmt.Println(qiniu_service.GenToken())
}
